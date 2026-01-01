package routes

import (
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/middlewares"
	"fmt"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(cfg *config.Config, router *gin.Engine) {
	router.Run(fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort))
}

func SetupRoute(cfg *config.Config, db *gorm.DB, enforcer *casbin.Enforcer) {
	route := gin.Default()

	route.Use(middlewares.TimeoutMiddleware(time.Duration(config.AppConfig.ContextRequestTimeout)))

	NewStaticRoute(route)

	public := route.Group("/api")
	NewAuthRoutes(public, db)

	protected := route.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	protected.Use(middlewares.CasbinMiddleware(enforcer))

	NewUserRoutes(protected, db)
	NewClassRoutes(protected, db)
	NewAssessmentRoutes(protected, db)
	NewExamRoutes(protected, db)
	NewStudentRoutes(protected, db)

	Run(cfg, route)
}
