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
	optionRepo := repositories.NewOptionRepo(db)
	questionRepo := repositories.NewQuestionRepo(db)
	answerRepo := repositories.NewAnswerRepo(db)
	examSubmissionRepo := repositories.NewExamSubmissionRepo(db)

	studentExamService := services.NewStudentExamService(studentExamRepo,
		examRepo,
		optionRepo,
		questionRepo,
		answerRepo,
		examSubmissionRepo)

	studentExamController := controllers.NewStudentExamController(studentExamService)

	exams := route.Group("/exams")
	{
		exams.POST("/:exam_id/start", studentExamController.StartExam)
		exams.GET("/:exam_id/questions", studentExamController.GetQuestions)
		// exams.GET("/:exam_id", studentExamController.GetExams)
		exams.PUT("/:exam_id/questions/:question_id/answer", studentExamController.AnswerQuestion)
	}
}
