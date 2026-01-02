package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerRepo interface {
	FindByUserAndQuestion(ctx context.Context, examID uuid.UUID, questionID uint, userID uint) (*models.ExamUserAnswer, error)
	Create(ctx context.Context, answer *models.ExamUserAnswer) error
	UpdateAnswer(ctx context.Context, id uint, answerID uint) error
}

type answerRepoImpl struct {
	DB *gorm.DB
}

func NewAnswerRepo(db *gorm.DB) AnswerRepo {
	return &answerRepoImpl{DB: db}
}

func (r *answerRepoImpl) UpdateAnswer(ctx context.Context, id uint, answerID uint) error {
	return r.DB.WithContext(ctx).Model(&models.ExamUserAnswer{}).
		Where("id = ?", id).
		Update("answer_id", answerID).Error
}

func (r *answerRepoImpl) Create(ctx context.Context, answer *models.ExamUserAnswer) error {
	return r.DB.WithContext(ctx).Create(answer).Error
}

func (r *answerRepoImpl) FindByUserAndQuestion(ctx context.Context, examID uuid.UUID, questionID uint, userID uint) (*models.ExamUserAnswer, error) {
	var answer models.ExamUserAnswer
	err := r.DB.WithContext(ctx).
		First(&answer, "exam_id = ? AND exam_question_id = ? AND user_id = ?",
			examID, questionID, userID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &answer, err
}
