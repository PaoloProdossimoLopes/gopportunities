package handler

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func openingHandler(openingRoutes *gin.RouterGroup) {
	const root = "/"
	openingRoutes.GET("/:openingId", showSpecificOpening)
	openingRoutes.GET(root, listOpeningsHandler)
	openingRoutes.POST(root, createOpeningHander)
	openingRoutes.DELETE(root, deleteOpeningHander)
	openingRoutes.PUT(root, updateOpeningHander)
}

func listOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}

func showSpecificOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET Specific Opening",
	})
}

func createOpeningHander(context *gin.Context) {
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

	context.JSON(200, createOpeningRequest)
}

func deleteOpeningHander(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE Opening",
	})
}

func updateOpeningHander(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "UPDATE Opening",
	})
}
