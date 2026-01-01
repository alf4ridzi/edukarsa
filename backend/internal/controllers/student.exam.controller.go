package controllers

import (
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"errors"
	"log"
	"net/http"

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

func (c *StudentExamController) GetExams(ctx *gin.Context) {
	// examID, err := uuid.Parse(ctx.Param("id"))
	// if err != nil {
	// 	helpers.BadRequest(ctx, "exam id is not valid")
	// 	return
	// }

}

func (c *StudentExamController) GetQuestions(ctx *gin.Context) {
	examID, err := uuid.Parse(ctx.Param("id"))
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
