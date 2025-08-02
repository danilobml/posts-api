package main

import (
	"github.com/danilobml/posts-api/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	r.Group("/")
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetOnePost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
}
