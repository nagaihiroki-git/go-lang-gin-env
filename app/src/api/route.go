package api

import (
	"github.com/gin-gonic/gin"
	"go-sample-env/pkg/controllers"
)

func DefineRoutes(r gin.IRouter) {
	v1 := r.Group("/")
	{
		v1.GET("/post", controllers.FetchPost())
		v1.POST("/create/post", controllers.CreatePost())
	}
}
