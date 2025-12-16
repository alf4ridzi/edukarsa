package routes

import (
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoutes(route *gin.RouterGroup, cfg *config.Config, db *gorm.DB) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, cfg)

	auth := route.Group("/auth")
	{
		auth.POST("/login", userController.Login)
		auth.POST("/register", userController.Register)
	}
}
