package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(ctx context.Context, user models.User) error
	ExistByUsername(ctx context.Context, username string) (bool, error)
	ExistByEmail(ctx context.Context, email string) (bool, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{DB: db}
}

func (r *userRepoImpl) Create(ctx context.Context, user models.User) error {
	return r.DB.WithContext(ctx).Create(&user).Error
}

func (r *userRepoImpl) ExistByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (r *userRepoImpl) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *userRepoImpl) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).First(&user, "username = ?", username).Error
	return &user, err
}

func (r *userRepoImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).First(&user, "email = ?", email).Error
	return &user, err
}
