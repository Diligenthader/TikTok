package repository

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"encoding/json"
	"strconv"

	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// gorm.Model
	Id int64 `gorm:"column:user_id; primary_key;"`
	//gorm:"column:user_id; primary_key;" 是一个标签，它提供了关于这个字段的额外信息给GORM库：
	//column:user_id 表示在数据库中，这个字段对应的列的名字是 user_id。
	//primary_key; 表示这个字段是表的主键。
	Name            string `gorm:"column:user_name"`
	Password        string `gorm:"column:password"`
	Follow          int64  `gorm:"column:follow_count"`
	Follower        int64  `gorm:"column:follower_count"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFav        int64  `gorm:"column:total_favorited"`
	FavCount        int64  `gorm:"column:favorite_count"`
}

func (User) TableName() string {
	return "users"
}

// 检查该用户名是否已经存在

func UserNameIsExist(userName string) error {
	db := common.GetDB() //ToDo:在对数据库进行操作时，
	user := User{}
	err := db.Where("user_name = ?", userName).Find(&user).Error
	//ToDo:First和Find的区别：
	//1.Find方法会返回满足条件的所有记录，而First方法只返回满足条件的第一条记录。
	//2.当没有找到匹配记录时，Find方法通常会返回一个空的切片或列表，而First方法则会返回该类型的零值，并可能伴有一个错误信息，如“record not found”。
	if err == nil {
		return errors.New("username exist")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

// 创建用户

func InsertUser(userName, password string) (*User, error) {
	db := common.GetDB()
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //这是系统自带的一个用于编码的函数.

	user := User{
		Name:            userName,
		Password:        string(hasedPassword),
		Follow:          0,
		Follower:        0,
		TotalFav:        0,
		FavCount:        0,
		Avatar:          "https://avatars.githubusercontent.com/u/128824086?v=3",
		BackgroundImage: "https://go.dev/images/gophers/ladder.svg",
		Signature:       "test sign",
	}
	result := db.Create(&user) //首先创建user这个结构体，同时也去初始化其中的一些量值，然后通过数据操作对相应的表进行数据的插入
	if result.Error != nil {
		return nil, result.Error
	}
	log.Infof("regist user:%+v", user)
	go CacheSetUser(user) //ToDo:redis的相关操作
	return &user, nil
}

//获取用户信息

func GetUserInfo(u interface{}) (User, error) {
	db := common.GetDB()
	user := User{}
	var err error
	switch u := u.(type) { //这是一个用于获取变量类型的一个操作
	case int64: //表示为接口传入一个int64即为用户的user_id
		user, err = CacheGetUser(u) //先从redis缓存中寻找是否有数据存储记录，如果有就可以直接将其提取出来，可以减缓数据库的压力
		if err == nil {             //表示从缓存中获取到了数据并且未报错，此时便可以直接返回相应的数据
			return user, nil
		}
		err = db.Where("user_id = ?", u).Find(&user).Error //如果在缓存中未找到数据项，那么将从数据库中Find则表示为查询所有符合该查询条件的数据

	case string: //表示为接口传入了一个string，即为用户的name
		err = db.Where("user_name = ?", u).Find(&user).Error
	default:
		err = errors.New("")
	}
	if err != nil {
		return user, errors.New("user error")
	}
	go CacheSetUser(user) //表示为第用户在第一次登录的时候没有发生错错误，此时比那可以在redis缓存中写入数据
	log.Infof("%+v", user)
	return user, nil
}

func CacheSetUser(u User) {
	uid := strconv.FormatInt(u.Id, 10) //u.Id表示为用户的主键
	err := common.CacheSet("user_"+uid, u)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}
}

func CacheGetUser(uid int64) (User, error) {
	key := strconv.FormatInt(uid, 10)
	data, err := common.CacheGet("user_" + key)
	user := User{}
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
