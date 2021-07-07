package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
