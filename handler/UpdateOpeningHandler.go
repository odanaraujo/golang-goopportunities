package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/gopportunities/handler/dto"
	"github.com/odanaraujo/gopportunities/handler/response"
	"github.com/odanaraujo/gopportunities/schemas"
	"net/http"
)

// @basePath /api/v1
// @Summary Update a opening
// @Description Update a opening
// @Tags openings
// @Accept  json
// @Produce  json
// @Param id path string true "Opening ID"
// @Param opening body dto.OpeningRequest{} true "Request body"
// @Success 200 {object} response.SuccessResponse
// Failure 400 {object} response.ErrorResponse
// Failure 500 {object} response.ErrorResponse
// @Router /opening/{id} [put]
func UpdateOpeningHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		logger.Errorf("id is empty")
		response.SendError(ctx, http.StatusBadRequest, "id is empty")
		return
	}

	request := dto.OpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("error on bind json: %s", err.Error())
		response.SendError(ctx, http.StatusBadRequest, "error on bind json")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("opening with id %s not found", id)
		response.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if opening.ID == 0 {
		logger.Errorf("opening with id %s not found", id)
		response.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	opening = request.ToOpening(opening)

	fmt.Println(opening)

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error on update opening with id %s", id)
		response.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error on update opening with id %s", id))
		return
	}

	response.SendSuccessOk(ctx, "opening-pdate", opening)

}
