package controllers

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"edukarsa-backend/internal/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExamController struct {
	service services.ExamService
}

func NewExamController(service services.ExamService) *ExamController {
	return &ExamController{service: service}
}

func (c *ExamController) GetQuestions(ctx *gin.Context) {
	examID, err := utils.ParseUUIDString(ctx.Param("id"))
	if err != nil {
		helpers.BadRequest(ctx, "EXAM ID IS NOT VALID")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	questions, err := c.service.ListExamQuestions(reqCtx, examID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "ujian tidak ditemukan", nil)
		case errors.Is(err, domain.ErrExamNotStarted):
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, err.Error(), nil)
		case errors.Is(err, domain.ErrExamAlreadyFinished):
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, err.Error(), nil)
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil mendapatkan pertanyaan ujian", questions)
}

func (c *ExamController) CreateQuestions(ctx *gin.Context) {
	var input []dto.CreateExamQuestionRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	examID, err := utils.ParseUUIDString(ctx.Param("id"))
	if err != nil {
		helpers.BadRequest(ctx, "invalid exam id")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	question, err := c.service.CreateQuestions(reqCtx, examID, input)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "ujian tidak ditemukan", nil)
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.ResponseJSON(ctx, http.StatusCreated, true, "berhasil membuat pertanyaan", question)
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
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "kelas tidak ada", nil)
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.ResponseJSON(ctx, http.StatusCreated, true, "berhasil membuat ujian", exam)
}
