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
// @Summary Delete a opening
// @Description Delete a opening
// @Tags openings
// @Accept  json
// @Produce  json
// @Param id path string true "Opening ID"
// @Success 204
// @Router /opening/{id} [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		logger.Infof("id is empty")
		response.SendError(ctx, http.StatusBadRequest, dto.ErrParamIsRequesd("id", "Param").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("opening with id %s not found", id)
		response.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("error on delete opening with id %s", id)
		response.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error on delete opening with id %s", id))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
