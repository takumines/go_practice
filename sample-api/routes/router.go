package routes

import (
	"github.com/takumines/go_practice/sample-api/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.GET("/posts", controllers.IndexPost)
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts/:post_id", controllers.ShowPost)
	router.PUT("/posts/:post_id", controllers.UpdatePost)
	router.DELETE("/posts/:post_id", controllers.DeletePost)

	posts := router.Group("/posts/:post_id")
	posts.GET("/comments", controllers.IndexComment)
	posts.POST("/comments", controllers.CreateComment)
	posts.GET("/comments/:comment_id", controllers.ShowComment)
	posts.PUT("/comments/:comment_id", controllers.UpdateComment)
	posts.DELETE("/comments/:comment_id", controllers.DeleteComment)
}
