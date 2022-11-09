package main

import (
	"github.com/gin-gonic/gin"
	"go-sample-env/api"
)

func main() {
	r := gin.Default()
	api.DefineRoutes(r)
	r.Run()

}
