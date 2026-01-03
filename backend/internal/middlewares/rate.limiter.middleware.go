package middlewares

import (
	"edukarsa-backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 5) // max 5 request/second

	return func(ctx *gin.Context) {
		if limiter.Allow() {
			ctx.Next()
		} else {
			helpers.ResponseJSON(ctx, http.StatusTooManyRequests, false, "Limit", nil)
			ctx.Abort()
			return
		}
	}
}
