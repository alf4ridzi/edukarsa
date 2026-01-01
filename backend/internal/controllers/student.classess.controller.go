package controllers

import (
	"context"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	exams, err := c.studentClassessService.GetExams(reqCtx, classID)
	if err != nil {
		switch {
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil mendapatkan list ujian", exams)
}
