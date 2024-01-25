package router

import (
	"github.com/gin-gonic/gin"

	"github.com/PaoloProdossimoLopes/gopportunities/docs"
	"github.com/PaoloProdossimoLopes/gopportunities/handler"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func attachRoutes(router *gin.Engine) {
	const basePath = "/api/v1/"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)

	handler.Handle(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
