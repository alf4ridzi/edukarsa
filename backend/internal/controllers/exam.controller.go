package controllers

import (
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"edukarsa-backend/internal/utils"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExamController struct {
	service services.ExamService
}

func NewExamController(service services.ExamService) *ExamController {
	return &ExamController{service: service}
}

func (c *ExamController) UpdateExam(ctx *gin.Context) {
	examID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		helpers.BadRequest(ctx, "exam id is not valid")
		return
	}

	var input dto.ExamUpdateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	exam, err := c.service.UpdateExam(ctx.Request.Context(), examID, input)
	if err != nil {
		switch {
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil memperbarui ujian", exam)
}

func (c *ExamController) GetQuestions(ctx *gin.Context) {
	examID, err := utils.ParseUUIDString(ctx.Param("id"))
	if err != nil {
		helpers.BadRequest(ctx, "EXAM ID IS NOT VALID")
		return
	}

	questions, err := c.service.ListExamQuestions(ctx.Request.Context(), examID)
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

	question, err := c.service.CreateQuestions(ctx.Request.Context(), examID, input)
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

	userID := ctx.GetUint64("user_id")

	exam, err := c.service.CreateNewExam(ctx.Request.Context(), uint(userID), input)
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
