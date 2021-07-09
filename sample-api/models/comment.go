package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func AllComment(post *Post, db *gorm.DB) (comments []Comment, err error) {
	rows, err := db.Where(&Comment{PostID: post.ID}).Find(&comments).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

func GetComment(id int, post *Post, db *gorm.DB) (comment Comment, err error) {
	comment = Comment{}
	err = db.Where(&Comment{PostID: post.ID}).First(&comment, id).Error
	return
}
