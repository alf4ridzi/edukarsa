package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExamRepo interface {
	Create(ctx context.Context, exam *models.Exam) error
	FindExamByID(ctx context.Context, id uuid.UUID) (*models.Exam, error)
	CreateOption(ctx context.Context, tx *gorm.DB, option *models.ExamOption) error
	CreateQuestion(ctx context.Context, tx *gorm.DB, question *models.ExamQuestion) error
	UpdateQuestion(ctx context.Context, tx *gorm.DB, question *models.ExamQuestion) error
}

type examRepoImpl struct {
	DB *gorm.DB
}

func NewExamRepo(db *gorm.DB) ExamRepo {
	return &examRepoImpl{DB: db}
}

func (r *examRepoImpl) FindExamByID(ctx context.Context, id uuid.UUID) (*models.Exam, error) {
	var exam models.Exam
	err := r.DB.WithContext(ctx).First(&exam, "id = ?", id).Error
	return &exam, err
}

func (r *examRepoImpl) CreateOption(ctx context.Context, tx *gorm.DB, option *models.ExamOption) error {
	return tx.WithContext(ctx).Create(option).Error
}

func (r *examRepoImpl) CreateQuestion(ctx context.Context, tx *gorm.DB, question *models.ExamQuestion) error {
	return tx.WithContext(ctx).Create(question).Error
}

func (r *examRepoImpl) UpdateQuestion(ctx context.Context, tx *gorm.DB, question *models.ExamQuestion) error {
	return tx.WithContext(ctx).Updates(question).Error
}

func (r *examRepoImpl) Create(ctx context.Context, exam *models.Exam) error {
	return r.DB.WithContext(ctx).Create(exam).Error
}
