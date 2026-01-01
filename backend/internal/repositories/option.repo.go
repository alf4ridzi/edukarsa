package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type OptionRepo interface {
	FindOptionByID(ctx context.Context, id uint) (*models.ExamOption, error)
}

type optionRepoImpl struct {
	DB *gorm.DB
}

func NewOptionRepo(db *gorm.DB) OptionRepo {
	return &optionRepoImpl{DB: db}
}

func (r *optionRepoImpl) FindOptionByID(ctx context.Context, id uint) (*models.ExamOption, error) {
	var option models.ExamOption
	err := r.DB.WithContext(ctx).First(&option, "id = ?", id).Error
	return &option, err
}
