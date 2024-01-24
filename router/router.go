package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	ginRouter := gin.Default()

	attachRoutes(ginRouter)

	ginRouter.Run(":8080" /*Optional*/)
}
