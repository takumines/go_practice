package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func gormConnect() *gorm.DB {
	dsn := "host=127.0.0.1 user=test password=test1234 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

func Posts() (posts []Post, err error) {
	db := gormConnect()

	rows, err := db.Find(&posts).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

func GetPost(id int) (*Post, error) {
	db := gormConnect()
	post := Post{}
	err := db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func Create(content string, author string) (err error) {
	db := gormConnect()
	post := Post{Content: content, Author: author}
	err = db.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func Update(id int, content string, author string) (err error) {
	db := gormConnect()

	BeforePost := Post{}

	err = db.Model(&BeforePost).Where("id", id).Updates(map[string]interface{}{"content": content, "author": author}).Error
	if err != nil {
		return err
	}
	return nil
}

func Delete(id int) (err error) {
	db := gormConnect()
	err = db.Delete(&Post{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func main() {
	Create("cccc", "ssss")

	Delete(10)

	Update(3, "gggggg", "9999978")

	posts, _ := Posts()
	fmt.Println(posts)

	post, _ := GetPost(4)
	fmt.Println(*post)
}
