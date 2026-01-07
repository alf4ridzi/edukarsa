package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type ExamScoreRepo interface {
	Create(ctx context.Context, score *models.ExamScore) error
}

type examScoreRepoImpl struct {
	DB *gorm.DB
}

func NewExamScoreRepo(db *gorm.DB) ExamScoreRepo {
	return &examScoreRepoImpl{DB: db}
}

func (r *examScoreRepoImpl) Create(ctx context.Context, score *models.ExamScore) error {
	return r.DB.WithContext(ctx).Create(score).Error
}
