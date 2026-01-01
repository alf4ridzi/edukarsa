package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseJSON(ctx *gin.Context, statusCode int, status bool, message string, data any) {
	ctx.JSON(statusCode, gin.H{
		"status":  status,
		"data":    data,
		"message": message,
	})
}

func OK(ctx *gin.Context, message string, data any) {
	ResponseJSON(ctx, http.StatusOK, true, message, data)
}

func BadRequest(ctx *gin.Context, message string) {
	ResponseJSON(ctx, http.StatusBadRequest, false, message, nil)
}
func InternalServerError(ctx *gin.Context, message string) {
	ResponseJSON(ctx, http.StatusInternalServerError, false, message, nil)
}
