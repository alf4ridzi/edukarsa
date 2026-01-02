package controllers

import (
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentExamController struct {
	studentExamService services.StudentExamService
}

func NewStudentExamController(studentExamService services.StudentExamService) *StudentExamController {
	return &StudentExamController{studentExamService: studentExamService}
}

func (c *StudentExamController) StartExam(ctx *gin.Context) {
	examID, err := uuid.Parse(ctx.Param("exam_id"))
	if err != nil {
		helpers.BadRequest(ctx, "id ujian tidak valid")
		return
	}

	userID := uint(ctx.GetUint64("user_id"))

	err = c.studentExamService.StartExam(
		ctx.Request.Context(),
		examID,
		userID,
	)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrExamNotStarted):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, err.Error(), nil)
		case errors.Is(err, domain.ErrExamAlreadyFinished):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, err.Error(), nil)
		case errors.Is(err, domain.ErrAlreadyStartExam):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, err.Error(), nil)
		case errors.Is(err, domain.ErrUserExamNotStarted):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, err.Error(), nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil memulai ujian", nil)
}

// func (c *StudentExamController) GetExams(ctx *gin.Context) {
// 	// examID, err := uuid.Parse(ctx.Param("id"))
// 	// if err != nil {
// 	// 	helpers.BadRequest(ctx, "exam id is not valid")
// 	// 	return
// 	// }

// }

func (c *StudentExamController) AnswerQuestion(ctx *gin.Context) {
	var input dto.StudentAnswerRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	examID, err := uuid.Parse(ctx.Param("exam_id"))
	if err != nil {
		helpers.BadRequest(ctx, "exam id tidak valid")
		return
	}

	questionIDUint64, err := strconv.ParseUint(ctx.Param("question_id"), 10, 64)
	if err != nil {
		helpers.BadRequest(ctx, "question id tidak valid")
		return
	}

	questionID := uint(questionIDUint64)

	userIDUint64 := ctx.GetUint64("user_id")
	userID := uint(userIDUint64)

	err = c.studentExamService.AnswerQuestion(ctx.Request.Context(), examID, questionID, input, userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrOptionNotBelongToQuestion):
			helpers.ResponseJSON(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		case errors.Is(err, domain.ErrSameAnswerSubmitted):
			helpers.ResponseJSON(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		case errors.Is(err, domain.ErrQuestionNotBelongToExam):
			helpers.ResponseJSON(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		case errors.Is(err, domain.ErrUserExamNotStarted):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, err.Error(), nil)
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil menjawab soal", nil)
}

func (c *StudentExamController) GetQuestions(ctx *gin.Context) {
	examID, err := uuid.Parse(ctx.Param("exam_id"))
	if err != nil {
		helpers.BadRequest(ctx, "exam id is not valid")
		return
	}

	questions, err := c.studentExamService.ListQuestions(ctx, examID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrExamAlreadyFinished):
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, err.Error(), nil)
		case errors.Is(err, domain.ErrExamNotAccessible):
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, err.Error(), nil)
		case errors.Is(err, domain.ErrExamNotStarted):
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, err.Error(), nil)
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "ujian tidak ada", nil)
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil mendapatkan pertanyaan", questions)
}
