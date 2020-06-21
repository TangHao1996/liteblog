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
type Video struct {
	gorm.Model
	Key       string `gorm:"unique;not null"`
	UserId    uint
	Title     string `gorm:"type:varchar(200)"`
	CoverPath string
	Path      string
	Visit     int `gorm:"default:0"`
}

func QueryVideosByPage(page int) {

}
