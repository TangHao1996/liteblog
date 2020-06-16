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

//登录查询
func QueryUserByEmailAndPassword(email, password string) (user User, err error) {
	return user, db.Where("email = ? and password = ?", email, password).Take(&user).Error
}

//查询用户名
func QueryUserByName(name string) (user User, err error) {
	return user, db.Where("name = ?", name).Take(&user).Error
}

//查询邮箱
func QueryUserByEmail(email string) (user User, err error) {
	return user, db.Where("email = ?", email).Take(&user).Error
}

func CreateUser(u User) error {
	return db.Create(&u).Error
}

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
	db.AutoMigrate(&User{})
	//若当前没有用户记录
	var cnt int
	if err := db.Model(&User{}).Count(&cnt).Error; err == nil && cnt == 0 {
		db.Create(&User{
			Name:     "admin",
			Email:    "tanghao1996.seu@gmail.com",
			Password: "123456",
			Avatar:   "/static/images/info-img.png",
			Role:     0,
		})
	}
}
