package handler

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary shows opening
// @Description shows job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		const statusCode = http.StatusInternalServerError
		context.JSON(statusCode, gin.H{
			"error":       "Internal server error",
			"reason":      "Error listing openings",
			"status_code": statusCode,
		})
		return
	}

	const statusCode = http.StatusOK
	context.JSON(statusCode, gin.H{
		"openings": openings,
	})
}
