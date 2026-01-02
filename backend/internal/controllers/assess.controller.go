package controllers

import (
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/services"
	"edukarsa-backend/internal/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AssessmentController struct {
	service services.AssessmentService
}

func NewAssessmentController(service services.AssessmentService) *AssessmentController {
	return &AssessmentController{service: service}
}

func (c *AssessmentController) Create(ctx *gin.Context) {
}

func (c *AssessmentController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	parseID, err := utils.ParseUUIDString(id)
	if err != nil {
		helpers.BadRequest(ctx, "hell nah, thats not uuid bro")
		return
	}

	err = c.service.Delete(ctx.Request.Context(), parseID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helpers.ResponseJSON(ctx, http.StatusNotFound, false, "tugas tidak ada", nil)
		default:
			helpers.ResponseJSON(ctx, http.StatusInternalServerError, false, "internal server error", nil)
		}
		return
	}

	helpers.OK(ctx, "berhasil hapus tugas", nil)
}
