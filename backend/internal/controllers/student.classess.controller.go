package controllers

import (
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentClassessController struct {
	studentClassessService services.StudentClassessService
}

func NewStudentClassessController(studentClassessService services.StudentClassessService) *StudentClassessController {
	return &StudentClassessController{studentClassessService: studentClassessService}
}

func (c *StudentClassessController) GetExams(ctx *gin.Context) {
	classID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		helpers.BadRequest(ctx, "invalid uuid")
		return
	}

	exams, err := c.studentClassessService.GetExams(ctx.Request.Context(), classID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "kelas tidak ditemukan", nil)
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil mendapatkan list ujian", exams)
}
