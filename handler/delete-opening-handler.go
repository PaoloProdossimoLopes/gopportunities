package handler

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary delete opening
// @Description delete job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [delete]
func DeleteOpeningHander(context *gin.Context) {
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

	if err := db.First(&opening, id).Error; err != nil {
		const statusCode = http.StatusNotFound
		context.JSON(statusCode, gin.H{
			"error":       "Not found",
			"reason":      "resource no found",
			"status_code": statusCode,
		})
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		const statusCode = http.StatusInternalServerError
		context.JSON(statusCode, gin.H{
			"error":       "Internal server error",
			"reason":      "error deleting the resource",
			"status_code": statusCode,
		})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
