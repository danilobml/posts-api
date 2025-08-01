package main

import (
	"github.com/danilobml/posts-api/cmd/api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	registerRoutes(r)

	r.Run()
}
