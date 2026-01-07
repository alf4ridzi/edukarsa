package services

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/mapper"
	"edukarsa-backend/internal/repositories"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentExamService interface {
	ListQuestions(ctx context.Context, examID uuid.UUID) ([]dto.ExamQuestionStudentResponse, error)
	AnswerQuestion(
		ctx context.Context,
		examID uuid.UUID, questionID uint,
		input dto.StudentAnswerRequest,
		userID uint) error
	StartExam(ctx context.Context, examID uuid.UUID, userID uint) error
	SubmitExam(ctx context.Context, examID uuid.UUID, userID uint) error
}

type studentExamServiceImpl struct {
	DB                 *gorm.DB
	studentExamRepo    repositories.StudentExamRepo
	examRepo           repositories.ExamRepo
	optionRepo         repositories.OptionRepo
	questionRepo       repositories.QuestionRepo
	answerRepo         repositories.AnswerRepo
	examSubmissionRepo repositories.ExamSubmissionRepo
}

func NewStudentExamService(
	DB *gorm.DB,
	studentExamRepo repositories.StudentExamRepo,
	examRepo repositories.ExamRepo,
	optionRepo repositories.OptionRepo,
	questionRepo repositories.QuestionRepo,
	answerRepo repositories.AnswerRepo,
	examSubmissionRepo repositories.ExamSubmissionRepo) StudentExamService {
	return &studentExamServiceImpl{
		DB:                 DB,
		studentExamRepo:    studentExamRepo,
		examRepo:           examRepo,
		optionRepo:         optionRepo,
		questionRepo:       questionRepo,
		answerRepo:         answerRepo,
		examSubmissionRepo: examSubmissionRepo}
}

func (s *studentExamServiceImpl) SubmitExam(ctx context.Context, examID uuid.UUID, userID uint) error {
	exam, err := s.examRepo.FindExamByID(ctx, examID)
	if err != nil {
		return err
	}

	submission, err := s.examSubmissionRepo.FindByExamIDAndUserID(ctx, exam.ID, userID)
	if err != nil {
		return err
	}

	if submission == nil {
		return domain.ErrSubmissionNotStarted
	}

	switch submission.Status {
	case domain.SubmissionSubmitted:
		return domain.ErrExamAlreadySubmitted
	case domain.SubmissionExpired:
		return domain.ErrExamExpired
	case domain.SubmissionOngoing:
	default:
		return domain.ErrInvalidSubmissionStatus
	}

	now := time.Now().UTC()

	endTime := submission.StartAt.Add(time.Duration(exam.Duration) * time.Minute)

	if now.After(endTime) {
		return domain.ErrExamDurationExceeded
	}

	questions, err := s.questionRepo.FindQuestionsByExamID(ctx, exam.ID)
	if err != nil {
		return err
	}

	answers, err := s.answerRepo.FindByExamAndUser(ctx, exam.ID, userID)
	if err != nil {
		return err
	}

	result := helpers.CalculateScore(questions, answers)

	return s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		examSubmissionRepo := repositories.NewExamSubmissionRepo(tx)
		examScoreRepo := repositories.NewExamScoreRepo(tx)

		if err := helpers.CanStudentStartExam(exam); err != nil {
			if errors.Is(err, domain.ErrExamAlreadyFinished) {
				submission.Status = domain.SubmissionExpired
				if err := examSubmissionRepo.Update(ctx, submission); err != nil {
					return err
				}
			}
			return err
		}

		now := time.Now().UTC()

		submission.SubmittedAt = &now
		submission.Status = domain.SubmissionSubmitted

		err = examSubmissionRepo.Update(ctx, submission)
		if err != nil {
			return err
		}

		score := &models.ExamScore{
			ExamID:     exam.ID,
			UserID:     userID,
			Correct:    result.Correct,
			Wrong:      result.Wrong,
			UnAnswered: result.UnAnswered,
			Score:      result.Score,
			FinishedAt: now,
		}

		err = examScoreRepo.Create(ctx, score)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *studentExamServiceImpl) StartExam(ctx context.Context, examID uuid.UUID, userID uint) error {
	exam, err := s.examRepo.FindExamByID(ctx, examID)
	if err != nil {
		return err
	}

	exist, err := s.examSubmissionRepo.ExistByExamIDAndUserID(ctx, exam.ID, userID)
	if err != nil {
		return err
	}

	if exist {
		return domain.ErrAlreadyStartExam
	}

	err = helpers.CanStudentStartExam(exam)
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	submission := &models.ExamSubmission{
		ExamID:  exam.ID,
		UserID:  userID,
		StartAt: now,
	}

	return s.examSubmissionRepo.Create(ctx, submission)
}

func (s *studentExamServiceImpl) AnswerQuestion(
	ctx context.Context,
	examID uuid.UUID, questionID uint,
	input dto.StudentAnswerRequest,
	userID uint) error {

	return s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		examRepo := repositories.NewExamRepo(tx)
		examSubmissionRepo := repositories.NewExamSubmissionRepo(tx)
		questionRepo := repositories.NewQuestionRepo(tx)
		optionRepo := repositories.NewOptionRepo(tx)
		answerRepo := repositories.NewAnswerRepo(tx)

		exam, err := examRepo.FindExamByID(ctx, examID)
		if err != nil {
			return err
		}

		submission, err := examSubmissionRepo.FindByExamIDAndUserID(ctx, exam.ID, userID)
		if err != nil {
			return err
		}

		if submission.Status != domain.SubmissionOngoing {
			switch submission.Status {
			case domain.SubmissionSubmitted:
				return domain.ErrExamAlreadySubmitted
			case domain.SubmissionExpired:
				return domain.ErrExamExpired
			default:
				return domain.ErrInvalidSubmissionStatus
			}
		}

		if err := helpers.CanStudentStartExam(exam); err != nil {
			if errors.Is(err, domain.ErrExamAlreadyFinished) {
				submission.Status = domain.SubmissionExpired
				err = examSubmissionRepo.Update(ctx, submission)
				if err != nil {
					return err
				}
			}

			return err
		}

		question, err := questionRepo.FindQuestionByID(ctx, questionID)
		if err != nil {
			return err
		}

		if question.ExamID != exam.ID {
			return domain.ErrQuestionNotBelongToExam
		}

		option, err := optionRepo.FindOptionByID(ctx, input.OptionID)
		if err != nil {
			return err
		}

		if option.ExamQuestionID != question.ID {
			return domain.ErrOptionNotBelongToQuestion
		}

		existing, err := answerRepo.FindByUserAndQuestion(ctx, exam.ID, question.ID, userID)
		if err != nil {
			return err
		}

		if existing != nil {
			if existing.AnswerID == input.OptionID {
				return domain.ErrSameAnswerSubmitted
			}

			return answerRepo.UpdateAnswer(ctx, existing.ID, input.OptionID)

		}

		answer := &models.ExamUserAnswer{
			ExamID:         exam.ID,
			UserID:         userID,
			ExamQuestionID: question.ID,
			AnswerID:       input.OptionID,
		}

		return answerRepo.Create(ctx, answer)
	})
}

func (s *studentExamServiceImpl) ListQuestions(ctx context.Context, examID uuid.UUID) ([]dto.ExamQuestionStudentResponse, error) {
	exam, err := s.studentExamRepo.FindExamByID(ctx, examID)
	if err != nil {
		return nil, err
	}

	err = helpers.CanStudentStartExam(exam)
	if err != nil {
		return nil, err
	}

	questions, err := s.studentExamRepo.FindQuestionsByExamID(ctx, exam.ID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ExamQuestionStudentResponse, 0, len(questions))

	for _, q := range questions {
		responses = append(responses, mapper.ToStudentQuestionResponse(q))
	}

	return responses, nil
}
