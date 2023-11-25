package repository

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"TikTokLite/util"
	"encoding/json"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Title         string `gorm:"column:title;"`
	Author        User   `gorm:"foreignkey:AuthorId"` //在这里，AuthorId 字段被用作 Video 表与 User 表之间的关联。通过设置 gorm:"foreignkey:AuthorId" 标签，GORM 将 AuthorId 字段作为外键字段，并将其关联到 User 表的主键。
	/*
		ToDo:这个字段定义了一个外键
		AuthorId 字段被用作 Video 结构体的外键。外键是用于建立表之间关系的字段，它指向另一个表的主键，以表示两个表之间的关联
		在这里，AuthorId 字段被用作 Video 表与 User 表之间的关联。通过设置 gorm:"foreignkey:AuthorId" 标签，GORM 将 AuthorId 字段作为外键字段，并将其关联到 User 表的主键。
		这意味着 AuthorId 字段的值将与 User 表中的主键值相对应，从而建立了 Video 表与 User 表之间的关系。通过这种关系，可以轻松地在查询 Video 记录时获取关联的 User 信息，例如作者的姓名、头像等。
		在使用 GORM 进行查询时，当查询到一条 Video 记录时，GORM 可以根据外键关系自动地加载关联的 User 数据，并将其填充到 Author 字段中。这样，在查询 Video 对象时，可以方便地访问与之相关联的作者信息，而无需单独执行额外的查询操作
	*/
}

func (Video) TableName() string {
	return "videos"
}

func InsertVideo(authorid int64, playurl, coverurl, title string) error {
	video := Video{
		AuthorId:      authorid,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		PublishTime:   util.GetCurrentTime(),
		Title:         title,
	}
	db := common.GetDB()
	err := db.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func GetVideoList(AuthorId int64) ([]Video, error) {
	var videos []Video
	author, err := GetUserInfo(AuthorId)
	if err != nil {
		return videos, err
	}
	db := common.GetDB()
	err = db.Where("author_id = ?", AuthorId).Order("video_id DESC").Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}
	for i := range videos {
		videos[i].Author = author
	}
	return videos, nil
}

func GetVideoListByFeed(currentTime int64) ([]Video, error) {
	var videos []Video
	db := common.GetDB()
	err := db.Where("publish_time < ?", currentTime).Limit(20).Order("video_id DESC").Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}
	for i, v := range videos {
		author, err := GetUserInfo(v.AuthorId)
		CacheSetAuthor(v.Id, v.AuthorId)
		if err != nil {
			return videos, err
		}
		videos[i].Author = author
	}
	return videos, nil
}

func CacheSetAuthor(videoid, authorid int64) {
	key := strconv.FormatInt(videoid, 10)
	err := common.CacheHSet("video", key, authorid)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}
}

func CacheGetAuthor(videoid int64) (int64, error) {
	key := strconv.FormatInt(videoid, 10)
	data, err := common.CacheHGet("video", key)
	if err != nil {
		return 0, err
	}
	uid := int64(0)
	err = json.Unmarshal(data, &uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
