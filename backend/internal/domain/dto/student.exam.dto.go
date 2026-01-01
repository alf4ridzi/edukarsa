package dto

import (
	"time"

	"github.com/google/uuid"
)

type ExamQuestionStudentResponse struct {
	ID       uint                 `json:"id"`
	Question string               `json:"question"`
	Options  []ExamOptionResponse `json:"options"`
}

type StudentExamResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	StartAt time.Time `json:"start_at"`
}
