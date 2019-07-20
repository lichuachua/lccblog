package models

import (
	//"github.com/jinzhu/gorm"
	"fmt"
)

type Note struct {
	Model
	Key     string `gorm:"unique:not null"`
	UserID  int
	User    User
	Title   string `gorm:"type:varchar(200)"`
	Summary string `gorm:"type:varchar(800)"`
	Content string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func QueryNoteByKey(key string) (note Note, err error) {
	return note, db.Where("`key` = ?", key).Take(&note).Error
}
func QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.Where("`key` = ? and user_id=?", key, userid).Take(&note).Error
}
func QueryNoteByPage(title string, page, limit int) (note []*Note, err error) {
	//输出  like "%title%"
	return note, db.Where("title like ?", fmt.Sprintf("%%%s%%", title)).Order("updated_at desc").Offset((page - 1) * limit).Limit(limit).Find(&note).Error
}

func DeleteNoteByUserIdAndKey(key string, userid int) error {
	return db.Delete(&Note{}, "`key` = ? and user_id = ?", key, userid).Error
}

func QueryNoteCount(title string) (count int, err error) {
	return count, db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Count(&count).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}
