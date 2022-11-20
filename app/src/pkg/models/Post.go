package models

import (
	database "go-sample-env/pkg/infra"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

func FetchPost(id int) Post {
	database.Init()
	post := Post{}

	database.Db.Take(&post, id)

	return Post{
		Title: post.Title,
		Body:  post.Body,
	}
}

func CreatePost(title string, body string) {
	database.Init()
	println(title)
	post := Post{
		Title: title,
		Body:  body,
	}
	database.Db.Create(&post)

}
