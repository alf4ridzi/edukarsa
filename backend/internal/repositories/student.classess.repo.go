package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type StudentClassessRepo interface {
	ListExamsByClassID(ctx context.Context, classID uint) ([]models.Exam, error)
}

type studentClassessRepoImpl struct {
	DB *gorm.DB
}

func NewStudentClassessRepo(db *gorm.DB) StudentClassessRepo {
	return &studentClassessRepoImpl{DB: db}
}

func (r *studentClassessRepoImpl) ListExamsByClassID(ctx context.Context, classID uint) ([]models.Exam, error) {
	var exams []models.Exam
	err := r.DB.WithContext(ctx).Find(&exams, "class_id = ?", classID).Error
	return exams, err
}
