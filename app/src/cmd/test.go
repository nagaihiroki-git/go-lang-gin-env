package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sample-env/pkg/infra"
	"gorm.io/gorm"
)

type Model struct {
	ID uint `gorm:"primary_key"`
}

type User struct {
	gorm.Model
	Name string
}

func Print(id int) gin.H {
	database.Init()
	user1 := User{}
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
