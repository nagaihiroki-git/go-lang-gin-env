package cmd

import "github.com/gin-gonic/gin"

func Print() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{})
	}
}
