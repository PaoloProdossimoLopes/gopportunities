package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	ginRouter := gin.Default()
	ginRouter.Run(":8080" /*Optional*/)
}
