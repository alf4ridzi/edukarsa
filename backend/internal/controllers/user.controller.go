package controllers

import (
	"context"
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"edukarsa-backend/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	service services.UserService
	cfg     *config.Config
}

func NewUserController(service services.UserService, cfg *config.Config) *UserController {
	return &UserController{service: service, cfg: cfg}
}

func (c *UserController) Login(ctx *gin.Context) {
	var reg models.Login
	if err := ctx.ShouldBindJSON(&reg); err != nil {
		helpers.BadRequest(ctx, err.Error())
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	user, err := c.service.Login(reqCtx, &reg)

	if err != nil {
		switch err {
		case models.ErrWrongPassword:
			helpers.ResponseJSON(ctx, http.StatusUnauthorized, false, "username/email/password salah", nil)
		case gorm.ErrRecordNotFound:
			helpers.ResponseJSON(ctx, http.StatusUnauthorized, false, "username/email/password salah", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	accessToken, err := utils.CreateAccessToken(user, c.cfg.AccessSecret, c.cfg.AccessTokenExpired)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	refreshToken, err := utils.CreateRefreshToken(user, c.cfg.RefreshSecret, c.cfg.RefreshTokenExpired)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	data := map[string]models.TokenResponse{
		"token": {
			Access:  accessToken,
			Refresh: refreshToken,
		},
	}

	helpers.OK(ctx, "berhasil login", data)
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

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
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
