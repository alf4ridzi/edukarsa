package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAssessmentRoutes(route *gin.RouterGroup, db *gorm.DB) {
	repos := repositories.NewAssessmentRepo(db)
	services := services.NewAssessmentService(repos)
	controllers := controllers.NewAssessmentController(services)

	assess := route.Group("/assessments")
	{
		assess.DELETE("/:id", controllers.Delete)
	}
}
