package models

import (
	"gorm.io/gorm"
)

type Post struct {
	// gorm.Modelはあまり使わない
	gorm.Model
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

// AllPost or FindAllPost
func AllPost(db *gorm.DB) (posts []Post, err error) {
	rows, err := db.Find(&posts).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

// GetPost or FindPostOfId
func GetPost(id int, db *gorm.DB) (post Post, err error) {
	post = Post{}
	err = db.Preload("Comments").First(&post, id).Error
	return
}
