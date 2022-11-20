package controllers

import (
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
