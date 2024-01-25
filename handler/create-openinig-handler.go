package handler

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary create opening
// @Description create new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [post]
func CreateOpeningHander(context *gin.Context) {
	createOpeningRequest := CreateOpeningRequest{}

	context.BindJSON(&createOpeningRequest)
	if validationCreateOpeningRequestError := createOpeningRequest.Validate(); validationCreateOpeningRequestError != nil {
		validationErrorMessage := validationCreateOpeningRequestError.Error()
		logger.Errorf("received invalid request to create opening job opportunity: %w", validationErrorMessage)
		const statusCode = http.StatusBadRequest
		context.JSON(statusCode, gin.H{
			"error":       "Bad request",
			"reason":      validationErrorMessage,
			"status_code": statusCode,
		})
		return
	}

	logger.Infof("request received: %+v", createOpeningRequest)

	opening := schemas.Opening{
		Role:     createOpeningRequest.Role,
		Link:     createOpeningRequest.Link,
		Location: createOpeningRequest.Location,
		Remote:   *createOpeningRequest.Remote,
		Salary:   createOpeningRequest.Salary,
	}
	if createDatabaseError := db.Create(&opening).Error; createDatabaseError != nil {
		logger.Errorf("creating opening: %v", createDatabaseError.Error())
		const statusCode = http.StatusInternalServerError
		context.JSON(statusCode, gin.H{
			"error":       "Internal server error",
			"reason":      createDatabaseError.Error(),
			"status_code": statusCode,
		})
		return
	}

	context.JSON(http.StatusOK, createOpeningRequest)
}
