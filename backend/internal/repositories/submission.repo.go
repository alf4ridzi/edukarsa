package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubmissionRepo interface {
	Create(ctx context.Context, submission *models.AssessmentSubmission) error
	FindAllByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]models.AssessmentSubmission, error)
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

func (r *SubmissionRepoImpl) FindAllByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]models.AssessmentSubmission, error) {
	var submissions []models.AssessmentSubmission
	err := r.DB.WithContext(ctx).Preload("Assessment").Find(&submissions, "assessment_id = ?", assessmentID).Error
	return submissions, err
}
