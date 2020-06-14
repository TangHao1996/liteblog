package models

import (
	"os"

	"github.com/astaxie/beego/logs"

	"github.com/jinzhu/gorm"                  //据说比beego自带的orm好用
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
	db, err = gorm.Open("mysql", "root:th1996@tcp(127.0.0.1:3306)/liteblog?charset=utf8")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	logs.Info("database connected")
}
