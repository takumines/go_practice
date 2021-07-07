package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

func All(db *gorm.DB) (posts []Post, err error) {
	rows, err := db.Find(&posts).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

func Retrieve(id int, db *gorm.DB) (post Post, err error) {
	post = Post{}
	err = db.First(&post, id).Error
	return
}
