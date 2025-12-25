package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewClassRoutes(route *gin.RouterGroup, db *gorm.DB) {
	classRepo := repositories.NewClassRepo(db)
	classService := services.NewClassService(classRepo)
	classController := controllers.NewClassController(classService)

	class := route.Group("/classes")
	{
		class.POST("", classController.Create)
		class.GET("", classController.GetUserClasses)
		class.POST("/join", classController.JoinClass)
	}
}
