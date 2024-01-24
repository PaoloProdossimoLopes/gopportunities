package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func attachRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("/openings", showOpeningHandler)
	v1.GET("/openings/:openingId", showSpecificOpening)
	v1.POST("/openings", createOpeningHander)
	v1.DELETE("/openings", deleteOpeningHander)
	v1.PUT("/openings", updateOpeningHander)
}

func showOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}

func showSpecificOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET Specific Opening",
	})
}

func createOpeningHander(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "POST Opening",
	})
}

func deleteOpeningHander(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE Opening",
	})
}

func updateOpeningHander(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "UPDATE Opening",
	})
}
