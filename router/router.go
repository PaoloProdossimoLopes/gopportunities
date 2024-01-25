package router

import (
	"github.com/PaoloProdossimoLopes/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	handler.Init()

	ginRouter := gin.Default()

	attachRoutes(ginRouter)

	ginRouter.Run(":8080" /*Optional*/)
}
