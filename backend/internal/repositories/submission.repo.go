package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type SubmissionRepo interface {
	Create(ctx context.Context, submission *models.AssessmentSubmission) error
}

type SubmissionRepoImpl struct {
	DB *gorm.DB
}

func NewSubmissionRepo(db *gorm.DB) SubmissionRepo {
	return &SubmissionRepoImpl{DB: db}
}

func (r *SubmissionRepoImpl) Create(ctx context.Context, submission *models.AssessmentSubmission) error {
	return r.DB.WithContext(ctx).Create(submission).Error
}
