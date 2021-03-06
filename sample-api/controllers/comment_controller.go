package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takumines/go_practice/sample-api/db"
	"github.com/takumines/go_practice/sample-api/models"
)

func IndexComment(c *gin.Context) {
	// post_idをでpostのデータ取得せずに、post_idを元にcommentモデルを取得する
	// もしくはEager Loadを利用する
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err := models.CheckPost(id, db.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "post not found",
		})
		return
	}
	comments, _ := models.AllComment(id, db.DB)
	c.JSON(http.StatusOK, comments)
}

func ShowComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err := models.CheckPost(postId, db.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "post not found",
		})
		return
	}
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	comment, err := models.GetComment(postId, commentId, db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func CreateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	post, err := models.GetPost(id, db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	comment := models.Comment{PostID: post.ID}
	if err := c.Bind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = db.DB.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSONP(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if err := models.CheckPost(postId, db.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "post not found",
		})
		return
	}

	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	comment, err := models.GetComment(postId, commentId, db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err := c.Bind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = db.DB.Updates(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if err := models.CheckPost(postId, db.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "post not found",
		})
		return
	}

	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	comment, err := models.GetComment(postId, commentId, db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	err = db.DB.Delete(&comment).Error
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
