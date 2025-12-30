package services

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/utils"

	"github.com/google/uuid"
)

type SubmissionService interface {
	SubmitSubmission(ctx context.Context, assessmentID string, userID uint, fileURL string, input models.AssessmentSubmissionRequest) (*models.AssessmentSubmission, error)
	GetAllSubmissionByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]models.AssessmentSubmission, error)
}

type SubmissionServiceImpl struct {
	repo repositories.SubmissionRepo
}

func NewSubmissionService(repo repositories.SubmissionRepo) SubmissionService {
	return &SubmissionServiceImpl{repo: repo}
}

func (s *SubmissionServiceImpl) SubmitSubmission(ctx context.Context, assessmentID string, userID uint, fileURL string, input models.AssessmentSubmissionRequest) (*models.AssessmentSubmission, error) {
	err := utils.ValidateUpload(input.File)
	if err != nil {
		return nil, err
	}

	parseAssmentID, err := utils.ParseUUIDString(assessmentID)
	if err != nil {
		return nil, err
	}

	submission := models.AssessmentSubmission{
		FileUrl:      fileURL,
		Description:  *input.Description,
		UserID:       userID,
		AssessmentID: parseAssmentID,
	}

	err = s.repo.Create(ctx, &submission)
	if err != nil {
		return nil, err
	}

	return &submission, nil
}

func (s *SubmissionServiceImpl) GetAllSubmissionByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]models.AssessmentSubmission, error) {
	return s.repo.FindAllByAssessmentID(ctx, assessmentID)
}
