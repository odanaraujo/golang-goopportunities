package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	})
}
