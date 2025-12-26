package controllers

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
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

func (c *ClassController) JoinClass(ctx *gin.Context) {
	var input models.JoinClassRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		helpers.BadRequest(ctx, "bad request")
		return
	}

	userID := ctx.GetUint64("user_id")

	err := c.service.JoinClass(ctx, input.ClassCode, userID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrAlreadyJoinedClass):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, "sudah bergabung", nil)
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "kelas tidak ada", nil)
		default:
			log.Println(err)
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

	err := c.service.CreateNewClass(reqCtx, userID, role, input)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrForbidden):
			helpers.ResponseJSON(ctx, http.StatusForbidden, false, "forbidden", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil membuat kelas baru", nil)
}
