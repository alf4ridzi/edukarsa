package routes

import "github.com/gin-gonic/gin"

func NewStaticRoute(route *gin.Engine) {
	route.Static("/assets", "./assets")
}
