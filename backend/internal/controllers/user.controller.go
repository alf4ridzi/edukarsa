package controllers

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Login(ctx *gin.Context) {
	var reg models.Login
	if err := ctx.ShouldBindJSON(&reg); err != nil {
		helpers.BadRequest(ctx, err.Error())
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	err := c.service.Login(reqCtx, &reg)

	if err != nil {
		switch err {
		case models.ErrWrongPassword:
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "username/email/password salah", nil)
		case gorm.ErrRecordNotFound:
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "username/email/password salah", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil login", nil)
}

func (c *UserController) Register(ctx *gin.Context) {
	var reg models.RegisterUser
	if err := ctx.ShouldBindJSON(&reg); err != nil {
		helpers.ResponseJSON(ctx, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	if reg.Password != reg.ConfirmPassword {
		helpers.BadRequest(ctx, "password tidak sama")
		return
	}

	// maks 5 detik
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	err := c.service.Register(reqCtx, &reg)

	if err != nil {
		switch err {
		case models.ErrUsernameExist:
			helpers.ResponseJSON(ctx, http.StatusConflict, false, "username sudah ada", nil)
		case models.ErrEmailExist:
			helpers.ResponseJSON(ctx, http.StatusConflict, false, "email sudah ada", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.ResponseJSON(ctx, http.StatusOK, true, "berhasil register", nil)
}
