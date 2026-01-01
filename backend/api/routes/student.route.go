package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewStudentRoutes(route *gin.RouterGroup, db *gorm.DB) {
	student := route.Group("/student")

	NewStudentClassessRoutes(student, db)
	NewStudentExamRoutes(student, db)
}
