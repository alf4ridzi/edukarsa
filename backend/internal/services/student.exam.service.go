package services

import (
	"context"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/mapper"
	"edukarsa-backend/internal/repositories"

	"github.com/google/uuid"
)

type StudentExamService interface {
	ListQuestions(ctx context.Context, examID uuid.UUID) ([]dto.ExamQuestionStudentResponse, error)
}

type studentExamServiceImpl struct {
	studentExamRepo repositories.StudentExamRepo
}

func NewStudentExamService(studentExamRepo repositories.StudentExamRepo) StudentExamService {
	return &studentExamServiceImpl{studentExamRepo: studentExamRepo}
}

func (s *studentExamServiceImpl) ListQuestions(ctx context.Context, examID uuid.UUID) ([]dto.ExamQuestionStudentResponse, error) {
	exam, err := s.studentExamRepo.FindExamByID(ctx, examID)
	if err != nil {
		return nil, err
	}

	err = helpers.CanStudentStartExam(exam)
	if err != nil {
		return nil, err
	}

	questions, err := s.studentExamRepo.FindQuestionsByExamID(ctx, exam.ID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ExamQuestionStudentResponse, 0, len(questions))

	for _, q := range questions {
		responses = append(responses, mapper.ToStudentQuestionResponse(q))
	}

	return responses, nil
}
