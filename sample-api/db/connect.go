package db

import (
	"github.com/takumines/go_practice/sample-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=127.0.0.1 user=test password=test1234 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = connect

	err = connect.AutoMigrate(&models.Post{}, &models.Comment{})
	if err != nil {
		panic(err)
	}
}
