package api

import (
	"github.com/gin-gonic/gin"
	"go-sample-env/cmd"
	"strconv"
)

func DefineRoutes(r gin.IRouter) {
	v1 := r.Group("/")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "SUCCESS!!!!",
			})
		})
		v1.GET("/test2", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Query("id"))
			c.JSON(200, cmd.Print(id))
		})
	}
}
