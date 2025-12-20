package routes

import (
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(cfg *config.Config, router *gin.Engine) {
	router.Run(fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort))
}

func SetupRoute(cfg *config.Config, db *gorm.DB) {
	route := gin.Default()

	public := route.Group("/api")
	NewAuthRoutes(public, db)

	private := route.Group("/api")
	private.Use(middlewares.AuthMiddleware)

	NewUserRoutes(private, db)
	Run(cfg, route)
}
