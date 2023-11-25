package repository

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
)

const (
	add = int64(1)
	sub = int64(-1)
)

type Relation struct {
	// gorm.Model
	Id       int64 `gorm:"column:relation_id; primary_key;"`
	Follow   int64 `gorm:"column:follow_id"`
	Follower int64 `gorm:"column:follower_id"`
}

func (Relation) TableName() string {
	return "relations"
}

func FollowAction(userId, toUserId int64) error {
	db := common.GetDB()
	relation := Relation{
		Follow:   userId,   //关注者
		Follower: toUserId, //被关注者
	}

	err := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Find(&Relation{}).Error
	if err != gorm.ErrRecordNotFound { //这行代码表示没有出现未找到记录的错误则表明在数据库中已有关注记录,即此时在这个Relation表中是有该数据项存在的，那么就就说明该toUserId已经存在，即为被关注过
		return errors.New("you have followed this user")
	}
	err = db.Create(&relation).Error
	//否则在这关系表中插入该新加入的数据项
	if err != nil {
		return err
	}
	go CacheChangeUserCount(userId, add, "follow")
	go CacheChangeUserCount(toUserId, add, "follower")
	return nil
}

func UnFollowAction(userId, toUserId int64) error {
	db := common.GetDB()
	err := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Delete(&Relation{}).Error
	//这是一个删除操作，将数据库表中的关注记录删除，则表示为取消关注
	if err != nil {
		return err
	}
	log.Debug("unfollow update user cache")
	go CacheChangeUserCount(userId, sub, "follow")
	go CacheChangeUserCount(toUserId, sub, "follower")
	return nil
}

func GetFollowList(userId int64, usertype string) ([]User, error) { //返回User的切片类型
	db := common.GetDB()
	re := []Relation{} //声明了一个Relation类型的切片，用于存储返回的数据.
	// joinArg := "follower"
	// if usertype == "follower" {
	// 	joinArg = "follow"
	// }
	// err := db.Joins("left join relations on users.user_id = relations."+joinArg+"_id").
	// 	Where("relations."+usertype+"_id = ?", userId).Find(&list).Error
	err := db.Where("relations."+usertype+"_id = ?", userId).Find(&re).Error
	/*
		ToDo:对数据库操作的解释
		db.Where("relations."+usertype+"_id = ?", userId)：这一行代码是在数据库中查询满足某个条件的记录。Where函数接受一个查询条件和对应的参数。在这个例子中，查询条件是"relations."+usertype+"_id = ?"，参数是userId。usertype是一个变量，它的值可能是user或其他值。这个查询条件的意思是在relations表中查找usertype+"_id"字段的值等于userId的记录。
	*/
	if err == gorm.ErrRecordNotFound {
		return []User{}, nil
	} else if err != nil {
		return nil, err
	}
	list := make([]User, len(re))
	for i, r := range re {
		uid := r.Follow           //表示为进行关注的用户id
		if usertype == "follow" { //ToDo：这是一个统一的操作，将被关注与关注进行返回，如果当usertype为"follow"时，说明此时需返回"followed"，反之亦然
			uid = r.Follower
		}
		list[i], _ = GetUserInfo(uid)
	}
	return list, nil
}

func CacheChangeUserCount(userid, op int64, ftype string) {
	uid := strconv.FormatInt(userid, 10)
	mutex, _ := common.GetLock("user_" + uid)
	defer common.UnLock(mutex)
	user, err := CacheGetUser(userid)
	if err != nil {
		log.Infof("user:%v miss cache", userid)
		return
	}
	switch ftype {
	case "follow":
		user.Follow += op
	case "follower":
		user.Follower += op
	case "like":
		user.FavCount += op
	case "liked":
		user.TotalFav += op
	}
	CacheSetUser(user)
}
