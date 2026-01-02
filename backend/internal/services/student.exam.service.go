package services

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/helpers"
	"edukarsa-backend/internal/mapper"
	"edukarsa-backend/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type StudentExamService interface {
	ListQuestions(ctx context.Context, examID uuid.UUID) ([]dto.ExamQuestionStudentResponse, error)
	AnswerQuestion(
		ctx context.Context,
		examID uuid.UUID, questionID uint,
		input dto.StudentAnswerRequest,
		userID uint) error
	StartExam(ctx context.Context, examID uuid.UUID, userID uint) error
}

type studentExamServiceImpl struct {
	studentExamRepo    repositories.StudentExamRepo
	examRepo           repositories.ExamRepo
	optionRepo         repositories.OptionRepo
	questionRepo       repositories.QuestionRepo
	answerRepo         repositories.AnswerRepo
	examSubmissionRepo repositories.ExamSubmissionRepo
}

func NewStudentExamService(studentExamRepo repositories.StudentExamRepo,
	examRepo repositories.ExamRepo,
	optionRepo repositories.OptionRepo,
	questionRepo repositories.QuestionRepo,
	answerRepo repositories.AnswerRepo,
	examSubmissionRepo repositories.ExamSubmissionRepo) StudentExamService {
	return &studentExamServiceImpl{studentExamRepo: studentExamRepo,
		examRepo:           examRepo,
		optionRepo:         optionRepo,
		questionRepo:       questionRepo,
		answerRepo:         answerRepo,
		examSubmissionRepo: examSubmissionRepo}
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

	exam, err := s.examRepo.FindExamByID(ctx, examID)
	if err != nil {
		return err
	}

	err = helpers.CanStudentStartExam(exam)
	if err != nil {
		return err
	}

	started, err := s.examSubmissionRepo.ExistByExamIDAndUserID(ctx, exam.ID, userID)
	if err != nil {
		return err
	}

	if !started {
		return domain.ErrUserExamNotStarted
	}

	question, err := s.questionRepo.FindQuestionByID(ctx, questionID)
	if err != nil {
		return err
	}

	if question.ExamID != exam.ID {
		return domain.ErrQuestionNotBelongToExam
	}

	option, err := s.optionRepo.FindOptionByID(ctx, input.OptionID)
	if err != nil {
		return err
	}

	if option.ExamQuestionID != question.ID {
		return domain.ErrOptionNotBelongToQuestion
	}

	existing, err := s.answerRepo.FindByUserAndQuestion(ctx, exam.ID, question.ID, userID)
	if err != nil {
		return err
	}

	if existing != nil && existing.AnswerID == input.OptionID {
		return domain.ErrSameAnswerSubmitted
	}

	if existing == nil {
		answer := &models.ExamUserAnswer{
			ExamID:         exam.ID,
			UserID:         userID,
			ExamQuestionID: question.ID,
			AnswerID:       input.OptionID,
		}

		return s.answerRepo.Create(ctx, answer)
	}

	return s.answerRepo.UpdateAnswer(ctx, existing.ID, input.OptionID)
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
