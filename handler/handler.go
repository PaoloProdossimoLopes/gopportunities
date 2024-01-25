package handler

import (
	"github.com/PaoloProdossimoLopes/gopportunities/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Handle(api *gin.RouterGroup) {
	openingHandler(api.Group("/openings"))
}

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
