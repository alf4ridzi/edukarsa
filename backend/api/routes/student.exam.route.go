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
	examRepo := repositories.NewExamRepo(db)

	studentExamService := services.NewStudentExamService(studentExamRepo, examRepo)

	studentExamController := controllers.NewStudentExamController(studentExamService)

	exams := route.Group("/exams")
	{
		exams.GET("/:exam_id/questions", studentExamController.GetQuestions)
		exams.GET("/:exam_id", studentExamController.GetExams)
		exams.PUT("/exams/:exam_id/questions/:question_id/answer")
	}
}
