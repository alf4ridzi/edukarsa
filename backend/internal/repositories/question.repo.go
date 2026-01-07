package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionRepo interface {
	FindQuestionByID(ctx context.Context, id uint) (*models.ExamQuestion, error)
	FindQuestionsByExamID(ctx context.Context, examID uuid.UUID) ([]models.ExamQuestion, error)
}

type questionRepoImpl struct {
	DB *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) QuestionRepo {
	return &questionRepoImpl{DB: db}
}

func (r *questionRepoImpl) FindQuestionsByExamID(ctx context.Context, examID uuid.UUID) ([]models.ExamQuestion, error) {
	var questions []models.ExamQuestion
	err := r.DB.WithContext(ctx).Find(&questions, "exam_id = ?", examID).Error
	return questions, err
}

func (r *questionRepoImpl) FindQuestionByID(ctx context.Context, id uint) (*models.ExamQuestion, error) {
	var question models.ExamQuestion
	err := r.DB.WithContext(ctx).First(&question, "id = ?", id).Error
	return &question, err
}
