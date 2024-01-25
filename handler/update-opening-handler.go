package handler

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary update opening
// @Description update job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [put]
func UpdateOpeningHander(context *gin.Context) {
	request := UpdateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		const statusCode = http.StatusBadRequest
		context.JSON(statusCode, gin.H{
			"error":       "Bad request",
			"reason":      err.Error(),
			"status_code": statusCode,
		})
		return
	}

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

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if createDatabaseError := db.Save(&opening).Error; createDatabaseError != nil {
		logger.Errorf("creating opening: %v", createDatabaseError.Error())
		const statusCode = http.StatusInternalServerError
		context.JSON(statusCode, gin.H{
			"error":       "Internal server error",
			"reason":      createDatabaseError.Error(),
			"status_code": statusCode,
		})
		return
	}

	context.JSON(http.StatusOK, opening)
}

type ErrorResponse struct {
	Error      string `json:"error"`
	Reason     string `json:"reason"`
	StatusCode int    `json:"status_code"`
}
