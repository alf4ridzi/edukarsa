package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExamSubmissionRepo interface {
	Create(ctx context.Context, submission *models.ExamSubmission) error
	ExistByExamIDAndUserID(ctx context.Context, examID uuid.UUID, userID uint) (bool, error)
}

type examSubmissionRepoImpl struct {
	DB *gorm.DB
}

func NewExamSubmissionRepo(db *gorm.DB) ExamSubmissionRepo {
	return &examSubmissionRepoImpl{DB: db}
}

func (r *examSubmissionRepoImpl) ExistByExamIDAndUserID(ctx context.Context, examID uuid.UUID, userID uint) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).
		Model(&models.ExamSubmission{}).
		Where("exam_id = ? AND user_id = ?", examID, userID).
		Count(&count).Error

	return count > 0, err
}

func (r *examSubmissionRepoImpl) Create(ctx context.Context, submission *models.ExamSubmission) error {
	return r.DB.WithContext(ctx).Create(submission).Error
}
