package services

import (
	"context"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/repositories"
	"log"

	"github.com/google/uuid"
)

type StudentClassessService interface {
	GetExams(ctx context.Context, classPublicID uuid.UUID) ([]dto.StudentExamResponse, error)
}

type studentClassessServiceImpl struct {
	studentClassessRepo repositories.StudentClassessRepo
	classRepo           repositories.ClassRepo
}

func NewStudentClassessService(studentClassessRepo repositories.StudentClassessRepo, classRepo repositories.ClassRepo) StudentClassessService {
	return &studentClassessServiceImpl{
		studentClassessRepo: studentClassessRepo,
		classRepo:           classRepo,
	}
}

func (s *studentClassessServiceImpl) GetExams(ctx context.Context, classPublicID uuid.UUID) ([]dto.StudentExamResponse, error) {
	class, err := s.classRepo.FindByPublicID(ctx, classPublicID)
	if err != nil {
		return nil, err
	}

	exams, err := s.studentClassessRepo.ListExamsByClassID(ctx, class.ID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.StudentExamResponse, 0, len(exams))

	for _, e := range exams {
		log.Println(e)

		err = helpers.CanStudentGetExam(&e)
		if err != nil {
			continue
		}

		resp := dto.StudentExamResponse{
			ID:      e.ID,
			Name:    e.Name,
			StartAt: e.StartAt,
		}

		responses = append(responses, resp)
	}

	return responses, nil
}
