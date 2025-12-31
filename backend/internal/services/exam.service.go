package services

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExamService interface {
	CreateNewExam(ctx context.Context, userID uint, input dto.CreateExamRequest) (*models.Exam, error)
	CreateQuestions(ctx context.Context, examID uuid.UUID, input []dto.CreateExamQuestionRequest) ([]dto.ExamQuestionResponse, error)
}

type ExamServiceImpl struct {
	DB        *gorm.DB
	repo      repositories.ExamRepo
	classRepo repositories.ClassRepo
}

func NewExamService(db *gorm.DB, repo repositories.ExamRepo, classRepo repositories.ClassRepo) ExamService {
	return &ExamServiceImpl{DB: db, repo: repo, classRepo: classRepo}
}

func (s *ExamServiceImpl) CreateQuestions(ctx context.Context, examID uuid.UUID, input []dto.CreateExamQuestionRequest) ([]dto.ExamQuestionResponse, error) {

	exam, err := s.repo.FindExamByID(ctx, examID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ExamQuestionResponse, 0, len(input))

	for _, q := range input {
		if len(q.Options) < 2 {
			return nil, domain.ErrMinimumOption
		}
		if q.CorrectIndex < 0 || q.CorrectIndex >= len(q.Options) {
			return nil, domain.ErrInvalidCorrectIndex
		}
	}

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		for _, inputQuestion := range input {
			question := models.ExamQuestion{
				ExamID:      exam.ID,
				Question:    inputQuestion.Question,
				Explanation: inputQuestion.Explanation,
			}

			if err := s.repo.CreateQuestion(ctx, tx, &question); err != nil {
				return err
			}

			options := make([]models.ExamOption, len(inputQuestion.Options))

			for i, opt := range inputQuestion.Options {
				options[i] = models.ExamOption{
					ExamQuestionID: question.ID,
					Option:         opt,
				}

				if err := s.repo.CreateOption(ctx, tx, &options[i]); err != nil {
					return err
				}
			}

			question.AnswerID = &options[inputQuestion.CorrectIndex].ID
			if err := s.repo.UpdateQuestion(ctx, tx, &question); err != nil {
				return err
			}

			respOptions := make([]dto.ExamOptionResponse, len(options))

			for i, o := range options {
				respOptions[i] = dto.ExamOptionResponse{
					ID:   o.ID,
					Text: o.Option,
				}
			}

			responses = append(responses, dto.ExamQuestionResponse{
				ID:        question.ID,
				Question:  question.Question,
				Options:   respOptions,
				CreatedAt: question.CreatedAt,
			})
		}

	})

	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s *ExamServiceImpl) CreateNewExam(ctx context.Context, userID uint, input dto.CreateExamRequest) (*models.Exam, error) {
	class, err := s.classRepo.FindByPublicID(ctx, input.ClassID)
	if err != nil {
		return nil, err
	}

	exam := models.Exam{
		Name:     input.Name,
		StartAt:  input.StartAt,
		Duration: input.Duration,
		Status:   "draft",
		ClassID:  class.ID,
	}

	err = s.repo.Create(ctx, &exam)
	if err != nil {
		return nil, err
	}

	return &exam, err
}
