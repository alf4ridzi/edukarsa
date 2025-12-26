package services

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"
)

type ClassService interface {
	CreateNewClass(ctx context.Context, userID uint, role string, input models.CreateClassRequest) error
	GetUserClasses(ctx context.Context, userID uint64) ([]models.Class, error)
	JoinClass(ctx context.Context, classCode string, userID uint64) error
	LeaveClass(ctx context.Context, classCode string, userID uint64) error
}

type classServiceImpl struct {
	repo repositories.ClassRepo
}

func NewClassService(repo repositories.ClassRepo) ClassService {
	return &classServiceImpl{repo: repo}
}

func (s *classServiceImpl) LeaveClass(ctx context.Context, classCode string, userID uint64) error {
	class, err := s.repo.FindByClassCode(ctx, classCode)
	if err != nil {
		return err
	}

	joined, err := s.repo.IsUserJoined(ctx, class.ID, userID)
	if err != nil {
		return err
	}

	if !joined {
		return models.ErrNotJoinedClass
	}

	return s.repo.Delete(ctx, class.ID, userID)
}

func (s *classServiceImpl) JoinClass(ctx context.Context, classCode string, userID uint64) error {

	class, err := s.repo.FindByClassCode(ctx, classCode)
	if err != nil {
		return err
	}

	joined, err := s.repo.IsUserJoined(ctx, class.ID, userID)
	if err != nil {
		return err
	}

	if joined {
		return models.ErrAlreadyJoinedClass
	}

	classUser := models.ClassUser{
		UserID:  uint(userID),
		ClassID: class.ID,
	}

	return s.repo.JoinClass(ctx, &classUser)
}

func (s *classServiceImpl) GetUserClasses(ctx context.Context, userID uint64) ([]models.Class, error) {
	return s.repo.FindByUserID(ctx, userID)
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
