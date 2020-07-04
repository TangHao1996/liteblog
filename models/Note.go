package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

const indexPageBNoteCount = 3

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
	Like    int    `gorm:"default:0"`
}

//多个表时要加表名
func QueryNoteByKey(key string) (note Note, err error) {
	return note, db.Where("note.key = ?", key).Take(&note).Error
}

//按页得到一组note
func QueryNotesByPage(searchWord string, page, limit int) (notes []*Note, err error) {
	return notes, db.Where("title like ?", fmt.Sprintf("%%%s%%", searchWord)).Offset((page - 1) * limit).Limit(limit).Find(&notes).Error
}

func QueryNoteCount(searchWord string) (count int, err error) {
	return count, db.Where("title like ?", fmt.Sprintf("%%%s%%", searchWord)).Model(&Note{}).Count(&count).Error
}

func CreateNote(note *Note) error {
	return db.Create(note).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}

func QueryNoteVisitByKey(key string) (visit int64) {
	rediconn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("redis connenct error")
		return
	}

	_, err = rediconn.Do("ZSCORE", "zset_visit", key)
	//若还没添加则先添加
	if err == redis.ErrNil {
		rediconn.Do("ZADD", "zset_visit", 0, key)
	}
	visit, _ = redis.Int64(rediconn.Do("ZSCORE", "zset_visit", key))
	return
}

func IncrVisit(key string) {
	rediconn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("redis connenct error")
		return
	}
	defer rediconn.Close()
	rediconn.Do("zincrby", "zset_visit", 1, key)
}

func QueryNoteLikeByKey(key string) (like int64) {
	rediconn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("redis connenct error")
		return
	}
	defer rediconn.Close()
	_, err = rediconn.Do("ZSCORE", "zset_like", key)
	//若还没添加则先添加
	if err == redis.ErrNil {
		rediconn.Do("ZADD", "zset_like", 0, key)
	}
	like, _ = redis.Int64(rediconn.Do("ZSCORE", "zset_like", key))
	return
}

func IncrLike(key string) {
	rediconn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("redis connenct error")
		return
	}
	defer rediconn.Close()
	rediconn.Do("zincrby", "zset_like", 1, key)
}

func GetIndexPageNotes() (notes []Note) {
	rediconn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("redis connenct error")
		return
	}
	keys, _ := redis.Strings(rediconn.Do("LRANGE", "indexPageNotes", 0, -1))

	notes = make([]Note, 0, indexPageBNoteCount)
	for _, key := range keys {
		n, _ := QueryNoteByKey(key)
		notes = append(notes, n)
	}

	return notes
}

func RefreshIndexPageNote() {
	rediconn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("redis connenct error")
		return
	}
	defer rediconn.Close()
	keys, err := redis.Strings(rediconn.Do("zrange", "zset_visit", 0, indexPageBNoteCount-1))
	if err != nil {
		logs.Error("redis zrange zset_visit error")
		return
	}

	for _, key := range keys {
		rediconn.Do("LPUSH", "indexPageNotes", key)
	}
}
