package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/gopportunities/handler/response"
	"github.com/odanaraujo/gopportunities/schemas"
	"net/http"
)

// @basePath /api/v1
// @Summary List all openings
// @Description List all openings
// @Tags openings
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.SuccessResponse
// Failure 400 {object} response.ErrorResponse
// Failure 500 {object} response.ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error on find openings: %v", err)
		response.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error on find openings: %v", err))
		return
	}

	response.SendSuccessOk(ctx, "list oppenings", openings)
}
