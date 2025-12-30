package controllers

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/models"
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

type ClassController struct {
	service services.ClassService
}

func NewClassController(service services.ClassService) *ClassController {
	return &ClassController{service: service}
}

func (c *ClassController) GetAssessments(ctx *gin.Context) {
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	publicID := ctx.Param("id")

	parseUUID, err := utils.ParseUUIDString(publicID)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	assessments, err := c.service.ListAssessmentByPublicID(reqCtx, parseUUID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "kelas tidak ada", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil mendapatkan assessments", assessments)
}

func (c *ClassController) CreateNewAssessment(ctx *gin.Context) {
	publicID := ctx.Param("id")

	var input models.CreateAssessmentRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	parseUUID, err := utils.ParseUUIDString(publicID)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	assessment, err := c.service.CreateNewAssessment(reqCtx, parseUUID, &input)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "kelas tidak ditemukan", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil menambah assessment", assessment)
}

func (c *ClassController) LeaveClass(ctx *gin.Context) {
	classCode := ctx.Param("code")
	if classCode == "" {
		helpers.BadRequest(ctx, "bad response")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	userID := ctx.GetUint64("user_id")

	err := c.service.LeaveClass(reqCtx, classCode, userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotJoinedClass):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, err.Error(), nil)
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "kelas tidak ada", nil)
		case errors.Is(err, domain.ErrCreatorCantLeave):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, err.Error(), nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil keluar kelas", nil)
}

func (c *ClassController) JoinClass(ctx *gin.Context) {
	classCode := ctx.Param("code")
	if classCode == "" {
		helpers.BadRequest(ctx, "bad response")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	userID := ctx.GetUint64("user_id")

	err := c.service.JoinClass(reqCtx, classCode, userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrAlreadyJoinedClass):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, "sudah bergabung", nil)
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "kelas tidak ada", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}
		return
	}

	helpers.OK(ctx, "berhasil bergabung ke kelas", nil)
}

func (c *ClassController) GetUserClasses(ctx *gin.Context) {
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	userID := ctx.GetUint64("user_id")

	classes, err := c.service.GetUserClasses(reqCtx, userID)
	if err != nil {
		log.Println(err)
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	helpers.OK(ctx, "berhasil mendapatkan kelas", classes)
}

func (c *ClassController) Create(ctx *gin.Context) {
	var input models.CreateClassRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	userID64 := ctx.GetUint64("user_id")
	role := ctx.GetString("role")

	userID := uint(userID64)

	class, err := c.service.CreateNewClass(reqCtx, userID, role, input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrForbidden):
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, "forbidden", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil membuat kelas baru", class)
}
