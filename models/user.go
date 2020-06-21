package models

import (
	"github.com/jinzhu/gorm"
)

/*
gorm.Model:
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
*/

//字段tag
//字段必须大写开头
type User struct {
	gorm.Model
	Name     string `gorm:"unique_index"`
	Email    string `gorm:"unique_index"`
	Password string
	Avatar   string
	Role     int //0: administrator, 1: user
}

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
