package middlewares

import (
	"edukarsa-backend/internal/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "no authorization header",
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "no authorization header",
		})
		return
	}

	tokenStr := parts[1]

	claims, err := utils.ValidateAccessToken(tokenStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "invalid or expired token",
		})
		return
	}

	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "internal server error",
		})
		return
	}

	ctx.Set("user_id", userID)
	ctx.Set("role", claims.Role)
	ctx.Next()
}
