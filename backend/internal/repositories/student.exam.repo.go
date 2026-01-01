package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentExamRepo interface {
	FindExamByID(ctx context.Context, id uuid.UUID) (*models.Exam, error)
	FindQuestionsByExamID(ctx context.Context, id uuid.UUID) ([]models.ExamQuestion, error)
	ExistExamByID(ctx context.Context, id uuid.UUID) (bool, error)
}

type studentExamRepoImpl struct {
	DB *gorm.DB
}

func NewStudentExamRepo(db *gorm.DB) StudentExamRepo {
	return &studentExamRepoImpl{DB: db}
}

func (r *studentExamRepoImpl) ExistExamByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *studentExamRepoImpl) FindExamByID(ctx context.Context, id uuid.UUID) (*models.Exam, error) {
	var exam models.Exam
	err := r.DB.WithContext(ctx).First(&exam, "id = ?", id).Error
	return &exam, err
}

func (r *studentExamRepoImpl) FindQuestionsByExamID(ctx context.Context, id uuid.UUID) ([]models.ExamQuestion, error) {
	var questions []models.ExamQuestion
	err := r.DB.WithContext(ctx).Preload("Options").Find(&questions, "exam_id = ?", id).Error
	return questions, err
}
