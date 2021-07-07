package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func gormConnect() *gorm.DB {
	dsn := "host=127.0.0.1 user=test password=test1234 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

func retrieve(id int) (post Post, err error) {
	db := gormConnect()
	post = Post{}
	err = db.First(&post, id).Error
	return
}

func all() (posts []Post, err error) {
	db := gormConnect()
	rows, err := db.Find(&posts).Rows()
	if err != nil {
		return
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	db := gormConnect()
	err = db.Create(&post).Error
	if err != nil {
		return
	}
	return
}

func (post *Post) Update() (err error) {
	db := gormConnect()
	err = db.Model(&post).Updates(&post).Error
	return
}

func (post *Post) Delete() (err error) {
	db := gormConnect()
	err = db.Delete(&post).Error
	return
}
