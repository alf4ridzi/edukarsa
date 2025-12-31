package controllers

import (
	"context"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type ExamController struct {
	service services.ExamService
}

func NewExamController(service services.ExamService) *ExamController {
	return &ExamController{service: service}
}

func (c *ExamController) Create(ctx *gin.Context) {
	var input dto.CreateExamRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	userID := ctx.GetUint64("user_id")

	exam, err := c.service.CreateNewExam(reqCtx, uint(userID), input)
	if err != nil {
		switch {
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil membuat ujian baru", exam)
}
