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
	Role     int `gorm:"default:1"` //0: administrator, 1: user
}
