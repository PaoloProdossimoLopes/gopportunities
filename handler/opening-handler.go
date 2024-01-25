package handler

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func openingHandler(openingRoutes *gin.RouterGroup) {
	const root = "/"
	openingRoutes.GET("/get", showSpecificOpening)
	openingRoutes.GET(root, listOpeningsHandler)
	openingRoutes.POST(root, createOpeningHander)
	openingRoutes.DELETE(root, deleteOpeningHander)
	openingRoutes.PUT(root, updateOpeningHander)
}

func listOpeningsHandler(context *gin.Context) {
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

func showSpecificOpening(context *gin.Context) {
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

	context.JSON(http.StatusOK, createOpeningRequest)
}

func deleteOpeningHander(context *gin.Context) {
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

func updateOpeningHander(context *gin.Context) {
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
