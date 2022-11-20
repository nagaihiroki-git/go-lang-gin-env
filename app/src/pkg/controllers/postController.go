package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sample-env/pkg/models"
	"strconv"
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

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		title := c.PostForm("title")
		body := c.PostForm("body")
		fmt.Println(title)
		models.CreatePost(title, body)
		c.JSON(200, gin.H{
			"message": "SUCCESS",
		})
	}
}
