package models

import (
	"gorm.io/gorm"
)

type Post struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

func AllPost(db *gorm.DB) (posts []Post, err error) {
	rows, err := db.Find(&posts).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

func GetPost(id int, db *gorm.DB) (post Post, err error) {
	post = Post{}
	err = db.Preload("Comments").First(&post, id).Error
	return
}

func CheckPost(id int, db *gorm.DB) (err error) {
	post := Post{}
	err = db.Where("id = ?", id).First(&post).Error
	return
}
