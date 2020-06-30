package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Key     string `gorm:"unique_index;not null"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	User    User   `json:"user"`
	NoteKey string `sql:"index" json:"noteKey"`
	Praise  int    `gorm:"default:0" json:"praise"`
}

func CreateComment(comment *Comment) error {
	return db.Create(comment).Error
}

func QueryCommentByNote(nkey string) (comments []*Comment, err error) {
	return comments, db.Where("comment.note_key = ?", nkey).Find(&comments).Error
}
