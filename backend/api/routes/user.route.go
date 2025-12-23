package routes

import (
	"edukarsa-backend/internal/controllers"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoutes(route *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	user := route.Group("/users")
	{
		user.GET("/me", userController.GetUser)
		user.PATCH("/me", userController.UpdateUser)
	}
}
