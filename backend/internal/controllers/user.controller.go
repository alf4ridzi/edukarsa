package controllers

import (
	"edukarsa-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Login(ctx *gin.Context) {

}

func (c *UserController) Register(ctx *gin.Context) {

}
