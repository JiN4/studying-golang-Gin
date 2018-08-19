package controllers

import (
	"studying-golang-Gin/models"
)

// User is
type User struct {
}

// NewUser ...
func NewUser() User {
	return User{}
}

// Get ...
func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(n)
	return user //userの構造体を返す
}

// Create ...
func (c User) Create(name string) {
	repo := models.NewUserRepository()
	repo.Create(name)
}
