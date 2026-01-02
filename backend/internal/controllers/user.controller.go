package controllers

import (
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"edukarsa-backend/internal/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var input models.UpdateUserData
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	userID := ctx.GetUint64("user_id")

	err := c.service.UpdateUserData(ctx.Request.Context(), userID, input)

	var pgErr *pgconn.PgError

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "user tidak ditemukan", nil)
		case errors.As(err, &pgErr) && pgErr.Code == "23505":
			helpers.ResponseJSON(ctx, http.StatusConflict, false, "username/email tidak tersedia", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.OK(ctx, "berhasil memperbarui data", nil)
}

func (c *UserController) Refresh(ctx *gin.Context) {
	var req models.RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ResponseJSON(ctx, http.StatusBadRequest, false, "refresh token required", nil)
		return
	}

	claims, err := utils.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		helpers.InternalServerError(ctx, "token invalid/expired")
		return
	}

	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusUnauthorized, false, "invalid token subject", nil)
		return
	}

	user, err := c.service.FindByID(ctx.Request.Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "user tidak ditemukan", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

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

	user, err := c.service.Login(ctx.Request.Context(), &reg)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrWrongPassword):
			helpers.ResponseJSON(ctx, http.StatusUnauthorized, false, "username/email/password salah", nil)
		case errors.Is(err, gorm.ErrRecordNotFound):
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

	data := models.TokenResponse{
		Access:  accessToken,
		Refresh: refreshToken,
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

	err := c.service.Register(ctx.Request.Context(), &reg)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUsernameExist):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, "username sudah ada", nil)
		case errors.Is(err, domain.ErrEmailExist):
			helpers.ResponseJSON(ctx, http.StatusConflict, false, "email sudah ada", nil)
		default:
			helpers.InternalServerError(ctx, "internal server error")
		}

		return
	}

	helpers.ResponseJSON(ctx, http.StatusOK, true, "berhasil register", nil)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")

	user, err := c.service.FindByID(ctx.Request.Context(), userID)
	if err != nil {
		helpers.InternalServerError(ctx, "internal server error")
		return
	}

	if user == nil {
		helpers.ResponseJSON(ctx, http.StatusNotFound, false, "user tidak ada", nil)
		return
	}

	data := map[string]*models.User{
		"user": user,
	}

	helpers.OK(ctx, "berhasil mendapatkan data user", data)
}
