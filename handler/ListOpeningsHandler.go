package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/gopportunities/handler/response"
	"github.com/odanaraujo/gopportunities/schemas"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error on find openings: %v", err)
		response.SendError(ctx, http.StatusBadRequest, fmt.Sprintf("Error on find openings: %v", err))
		return
	}

	response.SendSuccessOk(ctx, "list oppenings", openings)
}
