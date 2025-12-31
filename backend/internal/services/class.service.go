package services

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassService interface {
	CreateNewClass(ctx context.Context, userID uint, role string, input models.CreateClassRequest) (*models.Class, error)
	GetUserClasses(ctx context.Context, userID uint64) ([]models.Class, error)
	JoinClass(ctx context.Context, classCode string, userID uint64) error
	LeaveClass(ctx context.Context, classCode string, userID uint64) error
	CreateNewAssessment(ctx context.Context, publicID uuid.UUID, input *models.CreateAssessmentRequest) (*models.Assessment, error)
	ListAssessmentByPublicID(ctx context.Context, publicID uuid.UUID) ([]models.Assessment, error)
	ListExamsByClassPublicID(ctx context.Context, publicID uuid.UUID) ([]models.Exam, error)
}

type classServiceImpl struct {
	repo repositories.ClassRepo
}

func NewClassService(repo repositories.ClassRepo) ClassService {
	return &classServiceImpl{repo: repo}
}

func (s *classServiceImpl) ListExamsByClassPublicID(ctx context.Context, publicID uuid.UUID) ([]models.Exam, error) {
	class, err := s.repo.FindByPublicID(ctx, publicID)
	if err != nil {
		return nil, err
	}

	exams, err := s.repo.FindExamsByClassID(ctx, class.ID)
	return exams, err
}

func (s *classServiceImpl) ListAssessmentByPublicID(ctx context.Context, publicID uuid.UUID) ([]models.Assessment, error) {
	class, err := s.repo.FindByPublicID(ctx, publicID)
	if err != nil {
		return nil, err
	}

	return s.repo.FindAssessmentsByID(ctx, class.ID)
}

func (s *classServiceImpl) CreateNewAssessment(ctx context.Context, publicID uuid.UUID, input *models.CreateAssessmentRequest) (*models.Assessment, error) {
	class, err := s.repo.FindByPublicID(ctx, publicID)
	if err != nil {
		return nil, err
	}

	if class == nil {
		return nil, gorm.ErrRecordNotFound
	}

	assessment := models.Assessment{
		ClassID:    class.ID,
		Name:       input.Name,
		DeadlineAt: input.DeadlineAt,
	}

	err = s.repo.CreateForClass(ctx, &assessment)
	return &assessment, err
}

func (s *classServiceImpl) LeaveClass(ctx context.Context, classCode string, userID uint64) error {
	class, err := s.repo.FindByClassCode(ctx, classCode)
	if err != nil {
		return err
	}

	if userID == uint64(class.CreatedById) {
		return domain.ErrCreatorCantLeave
	}

	joined, err := s.repo.IsUserJoined(ctx, class.ID, userID)
	if err != nil {
		return err
	}

	if !joined {
		return domain.ErrNotJoinedClass
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
		return domain.ErrAlreadyJoinedClass
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

func (s *classServiceImpl) CreateNewClass(ctx context.Context, userID uint, role string, input models.CreateClassRequest) (*models.Class, error) {
	class := models.Class{
		Name:        input.Name,
		CreatedById: userID,
	}

	err := s.repo.CreateNewClass(ctx, &class)
	if err != nil {
		return nil, err
	}

	return &class, nil
}
