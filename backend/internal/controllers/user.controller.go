package controllers

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"edukarsa-backend/internal/utils"
	"net/http"
	"strconv"
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

func (c *UserController) Refresh(ctx *gin.Context) {
	var req models.RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ResponseJSON(ctx, http.StatusBadRequest, false, "refresh token required", nil)
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	claims, err := utils.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		helpers.InternalServerError(ctx, "token invalid/expired : "+err.Error())
		return
	}

	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusUnauthorized, false, "invalid token subject", nil)
		return
	}

	user, err := c.service.FindByID(reqCtx, userID)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	if user == nil {
		helpers.ResponseJSON(ctx, http.StatusUnauthorized, false, "user tidak ada", nil)
		return
	}

	accessToken, err := utils.CreateAccessToken(user.ID, user.Role.Name)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	data := map[string]string{
		"token": accessToken,
	}

	helpers.OK(ctx, "ok", data)
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

	accessToken, err := utils.CreateAccessToken(user.ID, user.Role.Name)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	refreshToken, err := utils.CreateRefreshToken(user.ID)
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

func (c *UserController) GetUser(ctx *gin.Context) {

}
