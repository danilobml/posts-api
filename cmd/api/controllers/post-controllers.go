package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/danilobml/posts-api/cmd/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var post *models.Post

	posts, err := post.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed fetching posts: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": posts})
}

func GetOnePost(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed fetching post: " + err.Error()})
		return
	}

	post := &models.Post{}
	post.ID = uint(intId)

	post, err = post.FindOne()
	if errors.Is(err, models.ErrPostNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed fetching post: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": post})
}

func CreatePost(c *gin.Context) {
	var requestBody struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed creating post: " + err.Error()})
		return
	}

	post := &models.Post{
		Title: requestBody.Title,
		Body:  requestBody.Body,
	}

	err = post.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed creating post: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": post})
}

func UpdatePost(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed deleting post: " + err.Error()})
		return
	}

	var requestBody struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed updating post: " + err.Error()})
		return
	}

	post := &models.Post{
		Title: requestBody.Title,
		Body:  requestBody.Body,
	}
	post.ID = uint(intId)

	err = post.Update()
	if errors.Is(err, models.ErrPostNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed updating post: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": post})
}

func DeletePost(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed deleting post: " + err.Error()})
		return
	}

	post := &models.Post{}
	post.ID = uint(intId)

	err = post.Delete()
	if errors.Is(err, models.ErrPostNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed deleting post: " + err.Error()})
		return
	}

	c.Status(204)
}
