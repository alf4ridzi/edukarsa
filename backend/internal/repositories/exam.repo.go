package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type ExamRepo interface {
	Create(ctx context.Context, exam *models.Exam) error
}

type ExamRepoImpl struct {
	DB *gorm.DB
}

func NewExamRepo(db *gorm.DB) ExamRepo {
	return &ExamRepoImpl{DB: db}
}

func (r *ExamRepoImpl) Create(ctx context.Context, exam *models.Exam) error {
	return r.DB.WithContext(ctx).Create(exam).Error
}
