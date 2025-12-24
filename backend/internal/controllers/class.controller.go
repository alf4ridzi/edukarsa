package controllers

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"time"

	"github.com/gin-gonic/gin"
)

type ClassController struct {
	service services.ClassService
}

func NewClassController(service services.ClassService) *ClassController {
	return &ClassController{service: service}
}

func (c *ClassController) Create(ctx *gin.Context) {
	var input models.CreateClassRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.BadRequest(ctx, "bad request")
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	err := c.service.CreateNewClass(reqCtx, input)
	if err != nil {
		switch {

		}
	}
}
