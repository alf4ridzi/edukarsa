package services

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"
)

type ClassService interface {
	CreateNewClass(ctx context.Context, userID uint, role string, input models.CreateClassRequest) error
}

type classServiceImpl struct {
	repo repositories.ClassRepo
}

func NewClassService(repo repositories.ClassRepo) ClassService {
	return &classServiceImpl{repo: repo}
}

func (s *classServiceImpl) CreateNewClass(ctx context.Context, userID uint, role string, input models.CreateClassRequest) error {
	if role != "teacher" {
		return models.ErrForbidden
	}

	class := models.Class{
		Name:        input.Name,
		CreatedById: userID,
	}

	return s.repo.CreateNewClass(ctx, &class)
}
