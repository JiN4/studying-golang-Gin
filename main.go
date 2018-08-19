package main

import (
	"github.com/JiN4/studying-golang-Gin/controllers"
	"github.com/gin-gonic/gin"
)

// main ...
func main() {
	router := gin.Default()

	router.GET("/:id", controllers.User.Get)
	router.POST("/insert", controllers.User.Create)

	router.Run(":8080")
}
