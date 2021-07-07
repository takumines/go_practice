package routes

import (
	"github.com/takumines/go_practice/sample-api/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.GET("/posts", controllers.IndexPost)
	router.GET("/posts/:id", controllers.ShowPost)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
}
