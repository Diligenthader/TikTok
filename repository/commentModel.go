package repository

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"encoding/json"
	"strconv"
	"time"
)

type Comment struct {
	CommentId int64  `gorm:"column:comment_id; primary_key;"` //完成Comment这个结构退和数据库对应的Comment表结构的映射
	UserId    int64  `gorm:"column:user_id"`
	VideoId   int64  `gorm:"column:video_id"`
	Comment   string `gorm:"column:comment"`
	Time      string `gorm:"column:time"`
}

func CommentAdd(userId, videoId int64, comment_text string) (*Comment, error) {
	db := common.GetDB()

	nowtime := time.Now().Format("01-02") //表示为获取当前的月份的表达式，同时也规定了相应的格式
	comment := Comment{
		UserId:  userId,
		VideoId: videoId,
		Comment: comment_text,
		Time:    nowtime,
	}
	result := db.Create(&comment)

	if result.Error != nil {
		return nil, result.Error
	}
	log.Infof("comment:%+v", comment)

	CacheDelCommentAll(videoId)

	return &comment, nil
}

func CommentDelete(videoId, comment_id int64) error { //ToDo:后续可以完善修改等操作，以及软删除的实现等.
	db := common.GetDB()
	commentTemp := Comment{}

	err := db.Where("comment_id = ?", comment_id).Take(&commentTemp).Error
	//.Take(&commentTemp)是一个操作，它获取查询结果的第一条记录，并将其存储在commentTemp变量中。&commentTemp表示commentTemp的地址，这样Take函数可以直接修改commentTemp的值。
	if err != nil {
		return err
	}
	CacheDelCommentAll(videoId)
	db.Delete(&commentTemp) //表示为将这个表删除.
	return nil

}

func CommentList(videoId int64) ([]Comment, error) {
	var comments []Comment //定义了一个comment切片，用于存储多个Comment类型的数据
	db := common.GetDB()
	var err error
	/* c := common.GetRE()
	values, _ := redis.Values(c.Do("lrange", videoId, "0", "-1"))
	for _,v := range values{

	} */

	comments, _ = CacheGetComment(videoId)
	log.Infof("comments-------------------------:%+v\n", comments)

	/* if err == nil {
		return comments, nil
	} */
	if comments != nil { //表示此时缓存不为空
		return comments, nil
	}

	err = db.Where("video_id = ?", videoId).Order("comment_id DESC").Find(&comments).Error

	CacheSetComment(videoId, comments)
	log.Infof("comments:%+v", comments)

	if err != nil {
		return nil, err
	}
	return comments, nil

}

// func GetVideoInfo(v interface{}) (*Video, error) {
// 	db := GetDB()
// 	video := Video{}
// 	err := db.Where("video_id = ?", v).Find(&video).Error
// 	if err != nil {
// 		return nil, errors.New("video error")
// 	}
// 	return &video, err
// }

// //根据user_id找到所有的用户信息
// func GetUser(v interface{}) (*User, error) {
// 	db := GetDB()
// 	var user User
// 	err := db.Where("user_id = ?", v).Find(&user).Error
// 	if err != nil {
// 		return nil, errors.New("user error")
// 	}
// 	return &user, nil

// }

/* func CacheSetComment(c Comment) {
	vid := strconv.FormatInt(c.VideoId, 10)
	err := common.CacheHSet("comment", vid, c)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}
} */

func CacheSetComment(videoId int64, c []Comment) {
	//for _, c1 := range c {
	//video_id := c1.VideoId

	vid := strconv.FormatInt(videoId, 10)
	err := common.CacheHSet("comment", vid, c)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}
}

//}

func CacheGetComment(vid int64) ([]Comment, error) {
	key := strconv.FormatInt(vid, 10)
	data, err := common.CacheHGet("comment", key)

	//var comments = make([]map[string]interface{},2)

	var comments []Comment
	//comment := Comment{}
	if err != nil {
		return comments, err
	}
	err = json.Unmarshal(data, &comments)
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func CacheDelCommentAll(videoId int64) {

	vid := strconv.FormatInt(videoId, 10) //将这个videoId这个字符串转化为int型.
	err := common.CacheDelHash("comment", vid)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}

}

/* func CacheDelCommentOne(videoId, comment_id int64) {

	vid := strconv.FormatInt(videoId, 10)
	cid := strconv.FormatInt(comment_id, 10)
	err := common.CacheDelHash2("comment", vid, cid)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}
} */
