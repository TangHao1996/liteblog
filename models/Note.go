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

type Note struct {
	gorm.Model
	Key     string `gorm:"unique_index;not null"`
	UserId  int
	User    User
	Title   string `gorm:"type:varchar(20)"`
	Summary string `gorm:"type:varchar(50)"`
	Content string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Like    int    `gorm:"default:0"`
}

//多个表时要加表名
func QueryNoteByKey(key string) (note Note, err error) {
	return note, db.Where("note.key = ?", key).Take(&note).Error
}

func CreateNote(note *Note) error {
	return db.Create(note).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}
