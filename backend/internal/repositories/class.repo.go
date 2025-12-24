package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type ClassRepo interface {
	Create(ctx context.Context, class *models.Class) error
	CreateNewClass(ctx context.Context, class *models.Class) error
}

type classRepoImpl struct {
	DB *gorm.DB
}

func NewClassRepo(db *gorm.DB) ClassRepo {
	return &classRepoImpl{DB: db}
}

func (r *classRepoImpl) Create(ctx context.Context, class *models.Class) error {
	return r.DB.WithContext(ctx).Create(class).Error
}

func (r *classRepoImpl) CreateNewClass(ctx context.Context, class *models.Class) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Updates({})
			return err
		}

		if err := tx.Create(&class).Error; err != nil {
			return err
		}

		return nil
	})
}
