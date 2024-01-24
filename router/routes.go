package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func attachRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "GET Opening",
		})
	})
}
