package services

import (
	"context"
	"edukarsa-backend/internal/repositories"

	"github.com/google/uuid"
)

type AssessmentService interface {
	Delete(ctx context.Context, id uuid.UUID) error
}

type AssessmentServiceImpl struct {
	repo repositories.AssessmentRepo
}

func NewAssessmentService(repo repositories.AssessmentRepo) AssessmentService {
	return &AssessmentServiceImpl{repo: repo}
}

func (s *AssessmentServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
