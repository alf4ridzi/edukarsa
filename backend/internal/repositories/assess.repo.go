package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssessmentRepo interface {
	Delete(ctx context.Context, id uuid.UUID) error
}

type AssessmentRepoImpl struct {
	DB *gorm.DB
}

func NewAssessmentRepo(db *gorm.DB) AssessmentRepo {
	return &AssessmentRepoImpl{DB: db}
}

func (r *AssessmentRepoImpl) Delete(ctx context.Context, id uuid.UUID) error {
	tx := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&models.Assessment{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
