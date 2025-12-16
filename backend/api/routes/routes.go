package routes

import (
	"edukarsa-backend/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(cfg *config.Config, router *gin.Engine) {
	router.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}

func SetupRoute(cfg *config.Config, db *gorm.DB) {
	route := gin.Default()

	public := route.Group("/api")
	NewAuthRoutes(public, cfg, db)

	Run(cfg, route)
}
