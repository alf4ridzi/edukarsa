package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewStudentClassessRoutes(route *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewStudentClassessRepo(db)
	classRepo := repositories.NewClassRepo(db)

	service := services.NewStudentClassessService(repo, classRepo)
	controller := controllers.NewStudentClassessController(service)

	classess := route.Group("/classess")
	{
		classess.GET("/:id/exams", controller.GetExams)
	}
}
