package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSubmissionRoute(assessments *gin.RouterGroup, api *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewSubmissionRepo(db)
	service := services.NewSubmissionService(repo)
	controller := controllers.NewSubmissionController(service)

	assessments.GET("/:id/submission", controller.GetSubmission)
	assessments.POST("/:id/submission", controller.Submission)

	submission := api.Group("/submissions")
	{
		submission.PATCH("/:id", controller.UpdateSubmission)
	}
}
