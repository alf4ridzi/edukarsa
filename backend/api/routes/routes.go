package routes

import (
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/middlewares"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(router *gin.Engine) {
	router.Run(fmt.Sprintf("%s:%d", config.AppConfig.ServerHost, config.AppConfig.ServerPort))
}

func SetupRoute(db *gorm.DB, enforcer *casbin.Enforcer) {
	route := gin.Default()

	route.Use(middlewares.TimeoutMiddleware(config.AppConfig.ContextRequestTimeout))

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

	Run(route)
}
