package api

import (
	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	v1 := r.Group("/")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "SUCCESS!!!!",
			})
		})
	}
}
