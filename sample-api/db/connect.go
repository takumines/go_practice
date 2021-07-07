package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=127.0.0.1 user=test password=test1234 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = connect
	if err != nil {
		panic(err.Error())
	}
}
