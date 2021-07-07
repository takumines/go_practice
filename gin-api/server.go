package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	router := gin.Default()

	router.GET("/posts", handleGetPosts)
	router.GET("/posts/:id", handleGet)
	router.POST("/posts", handlePost)
	router.PUT("/posts/:id", handlePut)
	router.DELETE("/posts/:id", handleDelete)

	router.Run(":8080")
}

func handleGetPosts(c *gin.Context) {
	posts, _ := all()

	c.JSON(http.StatusOK, posts)
}

func handleGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	post, err := retrieve(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func handlePost(c *gin.Context) {
	post := Post{}
	if err := c.Bind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err := post.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func handlePut(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	post, err := retrieve(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err := c.Bind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = post.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func handleDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	post, err := retrieve(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = post.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
