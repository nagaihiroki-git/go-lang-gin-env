package api

import (
	"go-sample-env/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	v1 := r.Group("/")
	{
		v1.GET("/post", controllers.FetchPost())
		v1.POST("/create/post", controllers.CreatePost())
	}
}
