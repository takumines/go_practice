package models

import "github.com/takumines/go_practice/sample-api/db"

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func All() (posts []Post, err error) {
	rows, err := db.DB.Find(&posts).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = db.DB.First(&post, id).Error
	return
}
