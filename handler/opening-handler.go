package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func openingHandler(openingRoutes *gin.RouterGroup) {
	const root = "/"
	openingRoutes.GET("/:openingId", showSpecificOpening)
	openingRoutes.GET(root, listOpeningsHandler)
	openingRoutes.POST(root, createOpeningHander)
	openingRoutes.DELETE(root, deleteOpeningHander)
	openingRoutes.PUT(root, updateOpeningHander)
}

func listOpeningsHandler(context *gin.Context) {
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
