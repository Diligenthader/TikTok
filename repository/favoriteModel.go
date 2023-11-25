package repository

import (
	"TikTokLite/common"
	"errors"

	"github.com/jinzhu/gorm"
)

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func LikeAction(uid, vid int64) error {
	db := common.GetDB()
	favorite := Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	err := db.Where("user_id = ? and video_id = ?", uid, vid).Find(&Favorite{}).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("you have liked this video")
	}
	//表示该条视频已被点赞收藏过，所以在数据库查询的时候会有该条记录.
	err = db.Create(&favorite).Error //否则，创建这条新的数据
	if err != nil {
		return err
	}
	authorid, _ := CacheGetAuthor(vid) //从缓存中获取视频的作者
	// go func() {
	// 	CacheChangeUserCount(uid, add, "like")
	// 	CacheChangeUserCount(authorid, add, "liked")
	// }()

	go CacheChangeUserCount(uid, add, "like")       //表示点赞视频，此时需要关联到当前的用户
	go CacheChangeUserCount(authorid, add, "liked") //表示为视频被点赞，即需要关联到视频的作者
	return nil
}

func UnLikeAction(uid, vid int64) error {
	db := common.GetDB()
	err := db.Where("user_id = ? and video_id = ?", uid, vid).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	authorid, _ := CacheGetAuthor(vid)
	// go func() {
	go CacheChangeUserCount(uid, sub, "like")
	go CacheChangeUserCount(authorid, sub, "liked")
	// }()
	return nil
}

func GetFavoriteList(uid int64) ([]Video, error) {
	var videos []Video //定义Video结构体的切片
	db := common.GetDB()
	err := db.Joins("left join favorites on videos.video_id = favorites.video_id").
		Where("favorites.user_id = ?", uid).Find(&videos).Error //此时的外键并未被填充，但由于videos与user表关联，所以在查询videos的时候也会加载外键关键的user

	/*
		ToDo:左连接查询
		左连接（Left Join）是数据库查询中的一种连接操作，它将两个表中的记录按照指定的条件进行连接，并返回左表中的所有记录，以及与之匹配的右表中的记录。
		在左连接中，左表是指在查询语句中位于左侧的表，而右表是指位于右侧的表。左连接以左表的记录为基础，将左表中的每一条记录与右表进行匹配。如果右表中存在与左表记录匹配的记录，那么左连接会返回左表记录和右表匹配记录的组合。如果右表中不存在匹配的记录，那么左连接仍然会返回左表记录，但是右表中的字段值会被填充为 NULL。

		这是一个左连接（LEFT JOIN）查询。在这个查询中，videos 是左表，favorites 是右表。查询会返回 videos 表中的所有行，以及 favorites 表中与 videos 表中的 video_id 匹配的行。如果在 favorites 表中没有与 videos 表相匹配的行，那么在结果集中会显示NULL值。
		此外，.Where("favorites.user_id = ?", uid) 这部分是在查询结果上应用了一个过滤条件，只返回 user_id 等于 uid 的记录
		db.Joins("left join favorites on videos.video_id = favorites.video_id").Where("favorites.user_id = ?", uid).Find(&videos) 来查询数据库。
		这个查询会联接 videos 和 favorites 表，查找 favorites.user_id 等于 uid 的所有视频，并将结果存储在 videos 切片中。
	*/

	if err == gorm.ErrRecordNotFound {
		return []Video{}, nil //说明此时并无记录,返回一个空的切片
	} else if err != nil {
		return nil, err
	}

	//ToDo:这是一个常见的处理切片存储多数据后提取的操作，即为遍历这个切片
	for i, v := range videos { //在使用range对videos进行遍历时，其会返回两个值，一个是元素的索引值，一个是其对应的元素值.
		author, err := GetUserInfo(v.AuthorId) //根据id查询视频的作者信息
		if err != nil {
			return videos, err
		}
		videos[i].Author = author
		//videos[i].Author=author这行代码将用户信息赋值给当前视频的Author字段。通过点赞收藏的信息得到user表中对应的信息.
	}

	return videos, nil
}
