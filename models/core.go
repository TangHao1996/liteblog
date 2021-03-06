package models

import (
	"os"

	"github.com/astaxie/beego/logs"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"                  //比beego自带的orm好用
	_ "github.com/jinzhu/gorm/dialects/mysql" //下划线 只执行其init函数
)

var (
	db *gorm.DB
)

func init() {
	var err error
	if err = os.MkdirAll("data", 0777); err != nil {
		panic(err.Error())
	}
	//parseTime是将时间转为time类型，不加会查询失败
	db, err = gorm.Open("mysql", "root:th1996@tcp(127.0.0.1:3306)/liteblog?charset=utf8&parseTime=true")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	db.SingularTable(true)
	logs.Info("database connected")

	//同步表结构, 如果没有表就创建一个
	db.AutoMigrate(&User{}, &Note{}, &Comment{})
	//若当前没有用户记录
	var cnt int
	if err := db.Model(&User{}).Count(&cnt).Error; err == nil && cnt == 0 {
		db.Create(&User{
			Name:     "admin",
			Email:    "tanghao1996.seu@gmail.com",
			Password: "th1996",
			Avatar:   "/static/images/info-img.png",
			Role:     0,
		})
	}

	initRedis()
}

func initRedis() {
	/*
		conn, err := cache.NewCache("redis", `{"key":"ViewsLikes","conn":":6379","dbNum":"0"}`)
		if err != nil {
			logs.Error("redis connenct error")
		}

		logs.Info("///////////////////redis connected/////////////////")
		conn.Put("test_key", "test_val", time.Second*300)
	*/
	rediconn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("redis connenct error")
		return
	}
	defer rediconn.Close()
	logs.Info("///////////////////redis connected/////////////////")
	//redis清空首页
	rediconn.Do("DEL", "indexPageNotes")

	//检查浏览量
	_, err = rediconn.Do("zrange", "zset_visit", 0, 10)
	if err != nil {
		logs.Error("无法读取浏览量")
	}

	RefreshIndexPageNote()
}
