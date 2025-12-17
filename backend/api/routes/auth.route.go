package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAuthRoutes(route *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	auth := route.Group("/auth")
	{
		auth.POST("/login", userController.Login)
		auth.POST("/register", userController.Register)
		auth.GET("/refresh", userController.Refresh)
	}
}
