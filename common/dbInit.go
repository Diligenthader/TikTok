// package repository
package common

import (
	"TikTokLite/config"
	"TikTokLite/log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DataBase *gorm.DB

func InitDatabase() {
	var err error
	conf := config.GetConfig()
	host := conf.Mysql.Host
	port := conf.Mysql.Port
	database := conf.Mysql.Database
	username := conf.Mysql.Username
	password := conf.Mysql.Password
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)
	//ToDo: 在对数据库进行操作时候args是我们所需要的一个格式
	log.Info(args)
	DataBase, err = gorm.Open("mysql", args) //ToDo：这是一个用于连接mysql数据库的操作，如果连接成功，则会返回一个指*gorm.DB的对象
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	log.Infof("connect database success,user:%s,database:%s", username, database)
}

func GetDB() *gorm.DB {
	return DataBase //获得这个对象，用于数据库的操作.
}
func CloseDataBase() {
	DataBase.Close()
}
