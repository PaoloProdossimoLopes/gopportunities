package router

import (
	"github.com/gin-gonic/gin"

	"github.com/PaoloProdossimoLopes/gopportunities/handler"
)

func attachRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	handler.Handle(v1)
}
