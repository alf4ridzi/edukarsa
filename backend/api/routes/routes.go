package routes

import (
	"edukarsa-backend/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(cfg *config.Config, router *gin.Engine) {
	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}

func SetupRoute(cfg *config.Config, db *gorm.DB) {
	route := gin.Default()

	public := route.Group("/api")
	NewUserRoutes(public, cfg, db)

	Run(cfg, route)
}
