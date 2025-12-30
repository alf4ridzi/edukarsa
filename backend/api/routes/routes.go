package routes

import (
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/middlewares"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(cfg *config.Config, router *gin.Engine) {
	router.Run(fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort))
}

func SetupRoute(cfg *config.Config, db *gorm.DB, enforcer *casbin.Enforcer) {
	route := gin.Default()

	NewStaticRoute(route)

	public := route.Group("/api")
	NewAuthRoutes(public, db)

	private := route.Group("/api")
	private.Use(middlewares.AuthMiddleware())
	private.Use(middlewares.CasbinMiddleware(enforcer))

	NewUserRoutes(private, db)
	NewClassRoutes(private, db)
	NewAssessmentRoutes(private, db)

	Run(cfg, route)
}
