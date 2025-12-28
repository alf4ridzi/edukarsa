package middlewares

import (
	"edukarsa-backend/internal/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		userID := ctx.GetUint64("user_id")

		userIDString := strconv.FormatUint(userID, 10)

		obj := ctx.FullPath()
		act := ctx.Request.Method

		allowed, err := e.Enforce(role, obj, act, userIDString)
		if err != nil {
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
			ctx.Abort()
			return
		}

		if !allowed {
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, "access denied", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}

}
