package controllers

import (
	"fmt"
	"strconv"

	"github.com/nagaihiroki-git/go-lang-gin-env/pkg/models"

	"github.com/gin-gonic/gin"
)

func FetchPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		Post := models.FetchPost(id)
		c.JSON(200, gin.H{
			"title": Post.Title,
			"body":  Post.Body,
		})
	}
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		body := c.PostForm("body")
		fmt.Println(title)
		models.CreatePost(title, body)
		c.JSON(200, gin.H{
			"message": "SUCCESS",
		})
	}
}
