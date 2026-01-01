package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewStudentExamRoutes(route *gin.RouterGroup, db *gorm.DB) {
	studentExamRepo := repositories.NewStudentExamRepo(db)
	studentExamService := services.NewStudentExamService(studentExamRepo)
	studentExamController := controllers.NewStudentExamController(studentExamService)

	exams := route.Group("/exams")
	{
		exams.GET("/:id/questions", studentExamController.GetQuestions)
		exams.GET("/:id", studentExamController.GetExams)
	}
}
