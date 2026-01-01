package services

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/mapper"
	"edukarsa-backend/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExamService interface {
	CreateNewExam(ctx context.Context, userID uint, input dto.CreateExamRequest) (*models.Exam, error)
	CreateQuestions(ctx context.Context, examID uuid.UUID, input []dto.CreateExamQuestionRequest) ([]dto.ExamQuestionResponse, error)
	ListExamQuestions(ctx context.Context, examID uuid.UUID) ([]dto.ExamQuestionTeacherResponse, error)
	UpdateExam(ctx context.Context, examID uuid.UUID, input dto.ExamUpdateRequest) (*models.Exam, error)
}

type examServiceImpl struct {
	DB        *gorm.DB
	repo      repositories.ExamRepo
	classRepo repositories.ClassRepo
}

func NewExamService(db *gorm.DB, repo repositories.ExamRepo, classRepo repositories.ClassRepo) ExamService {
	return &examServiceImpl{DB: db, repo: repo, classRepo: classRepo}
}

func (s *examServiceImpl) UpdateExam(ctx context.Context, examID uuid.UUID, input dto.ExamUpdateRequest) (*models.Exam, error) {
	exam, err := s.repo.FindExamByID(ctx, examID)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		exam.Name = *input.Name
	}

	if input.Status != nil {
		exam.Status = *input.Status
	}

	if input.Duration != nil {
		exam.Duration = *input.Duration
	}

	if input.StartAt != nil {
		exam.StartAt = *input.StartAt
	}

	if input.EndAt != nil {
		exam.EndAt = *input.EndAt
	}

	err = s.repo.Update(ctx, exam)
	return exam, err
}

func (s *examServiceImpl) ListExamQuestions(ctx context.Context, examID uuid.UUID) ([]dto.ExamQuestionTeacherResponse, error) {
	exam, err := s.repo.FindExamByID(ctx, examID)
	if err != nil {
		return nil, err
	}

	// now := time.Now().UTC()

	// if exam.StartAt.After(now) {
	// 	return nil, domain.ErrExamNotStarted
	// }

	// if now.After(exam.EndAt) {
	// 	return nil, domain.ErrExamAlreadyFinished
	// }

	questions, err := s.repo.ListQuestionsByExamID(ctx, exam.ID)
	responses := make([]dto.ExamQuestionTeacherResponse, 0, len(questions))

	for _, q := range questions {
		responses = append(responses, mapper.ToTeacherQuestionResponse(q))
	}

	return responses, nil
}

func (s *examServiceImpl) CreateQuestions(ctx context.Context, examID uuid.UUID, input []dto.CreateExamQuestionRequest) ([]dto.ExamQuestionResponse, error) {
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
				ID:              question.ID,
				Explanation:     question.Explanation,
				Question:        question.Question,
				Options:         respOptions,
				CorrectOptionID: *question.AnswerID,
				CreatedAt:       question.CreatedAt,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s *examServiceImpl) CreateNewExam(ctx context.Context, userID uint, input dto.CreateExamRequest) (*models.Exam, error) {
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
		EndAt:    input.EndAt,
	}

	err = s.repo.Create(ctx, &exam)
	if err != nil {
		return nil, err
	}

	return &exam, err
}
