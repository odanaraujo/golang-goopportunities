package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/gopportunities/handler/dto"
	"github.com/odanaraujo/gopportunities/handler/response"
	"github.com/odanaraujo/gopportunities/schemas"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	openingRequest := dto.OpeningRequest{}

	if err := ctx.BindJSON(&openingRequest); err != nil {
		logger.Errorf("Error binding opening: %v", err.Error())
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := openingRequest.Validate(); err != nil {
		logger.Errorf("Error validating opening: %v", err.Error())
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     openingRequest.Role,
		Company:  openingRequest.Company,
		Location: openingRequest.Location,
		Link:     openingRequest.Link,
		Remote:   *openingRequest.Remote,
		Salary:   openingRequest.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.SendSuccessOk(ctx, "create-opening", opening)
}
