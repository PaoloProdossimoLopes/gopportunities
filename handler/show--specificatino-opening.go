package handler

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary find opening
// @Description find job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} []OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/get [get]
func ShowSpecificOpening(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		const statusCode = http.StatusBadRequest
		context.JSON(statusCode, gin.H{
			"error":       "Bad request",
			"reason":      "id (type: string) is missing",
			"status_code": statusCode,
		})
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening).Error; err != nil {
		const statusCode = http.StatusNotFound
		context.JSON(statusCode, gin.H{
			"error":       "Not found",
			"reason":      "resource no found",
			"status_code": statusCode,
		})
		return
	}

	context.JSON(http.StatusOK, opening)
}
