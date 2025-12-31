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
	EvaluateSubmission(ctx context.Context, submissionID uuid.UUID, input models.EditSubmissionRequest) (*models.AssessmentSubmission, error)
}

type submissionServiceImpl struct {
	repo repositories.SubmissionRepo
}

func NewSubmissionService(repo repositories.SubmissionRepo) SubmissionService {
	return &submissionServiceImpl{repo: repo}
}

func (s *submissionServiceImpl) EvaluateSubmission(ctx context.Context, submissionID uuid.UUID, input models.EditSubmissionRequest) (*models.AssessmentSubmission, error) {
	submission, err := s.repo.FindByID(ctx, submissionID)
	if err != nil {
		return nil, err
	}

	if input.Score != nil {
		submission.Score = input.Score
	}

	if input.Feedback != nil {
		submission.Feedback = input.Feedback
	}

	err = s.repo.Update(ctx, submission)
	if err != nil {
		return nil, err
	}

	return submission, nil
}

func (s *submissionServiceImpl) SubmitSubmission(ctx context.Context, assessmentID string, userID uint, fileURL string, input models.AssessmentSubmissionRequest) (*models.AssessmentSubmission, error) {
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

func (s *submissionServiceImpl) GetAllSubmissionByAssessmentID(ctx context.Context, assessmentID uuid.UUID) ([]models.AssessmentSubmission, error) {
	return s.repo.FindAllByAssessmentID(ctx, assessmentID)
}
