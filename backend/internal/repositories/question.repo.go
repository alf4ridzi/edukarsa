package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type QuestionRepo interface {
	FindQuestionByID(ctx context.Context, id uint) (*models.ExamQuestion, error)
}

type questionRepoImpl struct {
	DB *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) QuestionRepo {
	return &questionRepoImpl{DB: db}
}

func (r *questionRepoImpl) FindQuestionByID(ctx context.Context, id uint) (*models.ExamQuestion, error) {
	var question models.ExamQuestion
	err := r.DB.WithContext(ctx).First(&question, "id = ?", id).Error
	return &question, err
}
