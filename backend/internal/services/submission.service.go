package services

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/utils"
)

type SubmissionService interface {
	SubmitSubmission(ctx context.Context, assessmentID string, userID uint, fileURL string, input models.AssessmentSubmissionRequest) error
}

type SubmissionServiceImpl struct {
	repo repositories.SubmissionRepo
}

func NewSubmissionService(repo repositories.SubmissionRepo) SubmissionService {
	return &SubmissionServiceImpl{repo: repo}
}

func (s *SubmissionServiceImpl) SubmitSubmission(ctx context.Context, assessmentID string, userID uint, fileURL string, input models.AssessmentSubmissionRequest) error {
	err := utils.ValidateUpload(input.File)
	if err != nil {
		return err
	}

	parseAssmentID, err := utils.ParseUUIDString(assessmentID)
	if err != nil {
		return err
	}

	submission := models.AssessmentSubmission{
		FileUrl:      fileURL,
		Description:  input.Description,
		UserID:       userID,
		AssessmentID: parseAssmentID,
	}

	return s.repo.Create(ctx, &submission)
}
