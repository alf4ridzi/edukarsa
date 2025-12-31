package mapper

import (
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/domain/models"
)

func ToStudentQuestionResponse(q models.ExamQuestion) dto.ExamQuestionStudentResponse {

	opts := make([]dto.ExamOptionResponse, 0, len(q.Options))

	for _, opt := range q.Options {
		opts = append(opts, dto.ExamOptionResponse{
			ID:   opt.ID,
			Text: opt.Option,
		})
	}

	return dto.ExamQuestionStudentResponse{
		ID:       q.ID,
		Question: q.Question,
		Options:  opts,
	}
}
