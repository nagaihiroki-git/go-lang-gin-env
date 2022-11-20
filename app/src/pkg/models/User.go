package models

import (
	"gorm.io/gorm"
)

type Model struct {
	ID uint `gorm:"primary_key"`
}

type User struct {
	gorm.Model
	Name string
}

//func Print(id int) gin.H {
//	database.Init()
//	user1 := models.User{}
//	print(id)
//	database.Db.Take(&user1)
//
//	fmt.Println(user1)
//
//	return gin.H{
//		"name": user1.Name,
//	}
//}
