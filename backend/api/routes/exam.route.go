package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewExamRoutes(route *gin.RouterGroup, db *gorm.DB) {
	examRepo := repositories.NewExamRepo(db)
	classRepo := repositories.NewClassRepo(db)

	examService := services.NewExamService(db, examRepo, classRepo)

	examController := controllers.NewExamController(examService)

	exams := route.Group("/exams")
	{
		exams.POST("", examController.Create)
		exams.POST("/:id/questions")
	}
}
