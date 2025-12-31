package services

import (
	"context"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"
)

type ExamService interface {
	CreateNewExam(ctx context.Context, userID uint, input dto.CreateExamRequest) (*models.Exam, error)
}

type ExamServiceImpl struct {
	repo      repositories.ExamRepo
	classRepo repositories.ClassRepo
}

func NewExamService(repo repositories.ExamRepo, classRepo repositories.ClassRepo) ExamService {
	return &ExamServiceImpl{repo: repo, classRepo: classRepo}
}

func (s *ExamServiceImpl) CreateNewExam(ctx context.Context, userID uint, input dto.CreateExamRequest) (*models.Exam, error) {
	class, err := s.classRepo.FindByPublicID(ctx, input.ClassID)
	if err != nil {
		return nil, err
	}

	exam := models.Exam{
		Name:     input.Name,
		StartAt:  input.StartAt,
		Duration: input.Duration,
		Status:   "draft",
		ClassID:  class.ID,
	}

	err = s.repo.Create(ctx, &exam)
	if err != nil {
		return nil, err
	}

	return &exam, err
}
