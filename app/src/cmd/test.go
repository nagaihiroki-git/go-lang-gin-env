package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	database "go-sample-env/pkg/infra"
	"go-sample-env/pkg/models"
)

func Print(id int) gin.H {
	database.Init()
	user1 := models.User{}
	print(id)
	database.Db.Take(&user1, id)

	fmt.Println(user1)

	return gin.H{
		"name": user1.Name,
	}
}

//
//type TodoJson struct {
//	ID     string `json:"id"`
//	Title  string `json:"title"`
//	Status string `js
//　　	on:"status"`
//}
//
//func Todo() gin.HandlerFunc {
//	var todos []TodoJson
//	return func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"todos": todos,

//		})
//	}
//}
