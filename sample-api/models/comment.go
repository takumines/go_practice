package models

import "gorm.io/gorm"

type Comment struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func AllComment(id int, db *gorm.DB) (comments []Comment, err error) {
	rows, err := db.Where("post_id = ?", id).Find(&comments).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

func GetComment(postId int, commentId int, db *gorm.DB) (comment Comment, err error) {
	comment = Comment{}
	err = db.Where("post_id = ? AND id = ?", postId, commentId).First(&comment).Error
	return
}
