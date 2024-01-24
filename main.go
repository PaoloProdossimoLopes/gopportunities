package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", get)

	router.Run()
}

func get(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
