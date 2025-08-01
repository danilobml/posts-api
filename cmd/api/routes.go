package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	r.Group("/")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, seu cuzao"})
	})
}
