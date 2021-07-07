package main

import (
	"go_practice/sample-api/db"
	"go_practice/sample-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	db.Init()
	routes.Setup(route)
	if err := route.Run(":8080"); err != nil {
		panic(err)
	}
}