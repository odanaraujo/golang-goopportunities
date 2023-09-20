package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/gopportunities/schemas"
	"net/http"
)

func SendError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func SendSuccessOk(ctx *gin.Context, op string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type SuccessResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
