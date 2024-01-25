package handler

import (
	"github.com/gin-gonic/gin"
)

func openingHandler(openingRoutes *gin.RouterGroup) {
	const root = "/"

	openingRoutes.GET("/get", ShowSpecificOpening)
	openingRoutes.GET(root, ListOpeningsHandler)
	openingRoutes.POST(root, CreateOpeningHander)
	openingRoutes.DELETE(root, DeleteOpeningHander)
	openingRoutes.PUT(root, UpdateOpeningHander)
}
