package main

import (
	"github.com/danilobml/posts-api/cmd/api/initializers"
	"github.com/danilobml/posts-api/cmd/api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
