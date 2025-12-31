package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateExamRequest struct {
	Name     string    `json:"name" binding:"required"`
	StartAt  time.Time `json:"start_at" binding:"required"`
	Duration int       `json:"duration" binding:"required"`
	ClassID  uuid.UUID `json:"class_id" binding:"required"`
}
