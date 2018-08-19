package controllers

import (
	"net/http"
	"strconv"

	"github.com/JiN4/studying-golang-Gin/models"
	"github.com/gin-gonic/gin"
)

var User = NewUser()

// User is
type user struct {
}

// NewUser ...
func NewUser() user {
	return user{}
}

// Get ...
func (u user) Get(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "id should be bigger than 0"})
		return
	}
	user := models.UserRepository.GetByID(id)
	// データを処理する
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Create ...
func (u user) Create(c *gin.Context) {
	name := c.PostForm("name")
	models.UserRepository.Create(name)
	c.Status(http.StatusCreated)
}
