package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/gopportunities/handler/dto"
	"github.com/odanaraujo/gopportunities/handler/response"
	"github.com/odanaraujo/gopportunities/schemas"
	"net/http"
)

// @basePath /api/v1
// @Summary get a opening
// @Description get a opening
// @Tags openings
// @Accept  json
// @Produce  json
// @Param id path string true "Opening ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /opening/{id} [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		logger.Infof("id not found")
		response.SendError(ctx, http.StatusBadRequest, dto.ErrParamIsRequesd("id", "Param").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.Find(&opening, id).Error; err != nil {
		logger.Errorf("error on find opening with id %s", id)
		response.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if opening.ID == 0 {
		logger.Infof("opening with id %s not found", id)
		response.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	response.SendSuccessOk(ctx, "show opening", opening)
}
