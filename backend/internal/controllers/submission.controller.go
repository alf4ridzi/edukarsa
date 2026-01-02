package controllers

import (
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"edukarsa-backend/internal/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubmissionController struct {
	service services.SubmissionService
}

func NewSubmissionController(service services.SubmissionService) *SubmissionController {
	return &SubmissionController{service: service}
}

func (c *SubmissionController) UpdateSubmission(ctx *gin.Context) {
	var input models.EditSubmissionRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	submissionIDString := ctx.Param("id")
	submissionID, err := utils.ParseUUIDString(submissionIDString)
	if err != nil {
		helpers.BadRequest(ctx, "NOT UUID")
		return
	}

	submission, err := c.service.EvaluateSubmission(ctx.Request.Context(), submissionID, input)
	if err != nil {
		switch {
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil evaluate tugas", submission)
}

func (c *SubmissionController) GetSubmission(ctx *gin.Context) {

	assessmentIDString := ctx.Param("id")
	assessmentID, err := utils.ParseUUIDString(assessmentIDString)

	if err != nil {
		helpers.BadRequest(ctx, "hell nah, thats not UUID bro")
		return
	}

	submission, err := c.service.GetAllSubmissionByAssessmentID(ctx.Request.Context(), assessmentID)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	helpers.OK(ctx, "berhasil", submission)
}

func (c *SubmissionController) Submission(ctx *gin.Context) {
	var input models.AssessmentSubmissionRequest

	if err := ctx.ShouldBind(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	assesmentID := ctx.Param("id")
	userID := ctx.GetUint64("user_id")

	ext := utils.GetExtension(input.File)
	fileName := fmt.Sprintf("%s%s", uuid.New(), ext)

	filePath := filepath.Join("assets", "images", "submissions", fileName)
	publicURL := "/" + path.Join("assets", "images", "submissions", fileName)

	submission, err := c.service.SubmitSubmission(ctx.Request.Context(), assesmentID, uint(userID), publicURL, input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrFileSizeTooBig):
			helpers.ResponseJSON(ctx, http.StatusRequestEntityTooLarge, false, "file terlalu besar", nil)
		case errors.Is(err, domain.ErrInvalidExtension):
			helpers.ResponseJSON(ctx, http.StatusUnsupportedMediaType, false, "file tidak disupport", nil)
		default:
			log.Println(err)
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	if err := ctx.SaveUploadedFile(input.File, filePath, 0755); err != nil {
		log.Println(err)
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	helpers.OK(ctx, "ok", submission)
}
