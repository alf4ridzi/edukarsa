package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSubmissionRoute(route *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewSubmissionRepo(db)
	service := services.NewSubmissionService(repo)
	controller := controllers.NewSubmissionController(service)

	route.GET("/:id/submission", controller.GetSubmission)
	route.POST("/:id/submission", controller.Submission)
}
