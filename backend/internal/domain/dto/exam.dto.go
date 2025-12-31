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

type CreateExamQuestionRequest struct {
	Question     string   `json:"question" binding:"required"`
	Explanation  *string  `json:"explanation"`
	Options      []string `json:"options" binding:"required,min=2"`
	CorrectIndex int      `json:"correct_index" binding:"required"`
}

type ExamQuestionResponse struct {
	ID              uint                 `json:"id"`
	Question        string               `json:"question" binding:"required"`
	Explanation     *string              `json:"explanation"`
	Options         []ExamOptionResponse `json:"options" binding:"required,min=2"`
	CorrectOptionID uint                 `json:"correct_option_id" binding:"required"`
	CreatedAt       time.Time            `json:"created_at"`
}

type ExamOptionResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}
