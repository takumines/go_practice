package routes

import (
	"github.com/takumines/go_practice/sample-api/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.GET("/posts", controllers.IndexPost)
	router.GET("/posts/:post_id", controllers.ShowPost)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:post_id", controllers.UpdatePost)
	router.DELETE("/posts/:post_id", controllers.DeletePost)

	router.GET("/posts/:post_id/comments", controllers.IndexComment)
	router.GET("/posts/:post_id/comments/:comment_id", controllers.ShowComment)
	router.POST("/posts/:post_id/comments", controllers.CreateComment)
}
