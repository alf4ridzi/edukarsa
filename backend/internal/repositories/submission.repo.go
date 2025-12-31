package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubmissionRepo interface {
	Create(ctx context.Context, submission *models.AssessmentSubmission) error
	Update(ctx context.Context, submission *models.AssessmentSubmission) error
	FindAllByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]models.AssessmentSubmission, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.AssessmentSubmission, error)
}

type submissionRepoImpl struct {
	DB *gorm.DB
}

func NewSubmissionRepo(db *gorm.DB) SubmissionRepo {
	return &submissionRepoImpl{DB: db}
}

func (r *submissionRepoImpl) FindByID(ctx context.Context, id uuid.UUID) (*models.AssessmentSubmission, error) {
	var submission models.AssessmentSubmission
	err := r.DB.WithContext(ctx).First(&submission, "id = ?", id).Error
	return &submission, err
}

func (r *submissionRepoImpl) Update(ctx context.Context, submission *models.AssessmentSubmission) error {
	return r.DB.WithContext(ctx).Model(submission).Updates(submission).Error
}

func (r *submissionRepoImpl) Create(ctx context.Context, submission *models.AssessmentSubmission) error {
	return r.DB.WithContext(ctx).Create(submission).Error
}

func (r *submissionRepoImpl) FindAllByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]models.AssessmentSubmission, error) {
	var submissions []models.AssessmentSubmission
	err := r.DB.WithContext(ctx).Preload("Assessment").Find(&submissions, "assessment_id = ?", assessmentID).Error
	return submissions, err
}
