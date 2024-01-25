package handler

import (
	"github.com/gin-gonic/gin"
)

func Handle(api *gin.RouterGroup) {
	openingHandler(api.Group("/openings"))
}
