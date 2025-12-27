package repositories

import (
	"context"
	"edukarsa-backend/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassRepo interface {
	Create(ctx context.Context, class *models.Class) error
	CreateNewClass(ctx context.Context, class *models.Class) error
	FindByUserID(ctx context.Context, userID uint64) ([]models.Class, error)
	JoinClass(ctx context.Context, classUser *models.ClassUser) error
	ExistByID(ctx context.Context, classID uint) (bool, error)
	IsUserJoined(ctx context.Context, classID uint, userID uint64) (bool, error)
	ExistByClassCode(ctx context.Context, classCode string) (bool, error)
	FindByClassCode(ctx context.Context, classCode string) (*models.Class, error)
	Delete(ctx context.Context, classID uint, userID uint64) error
	CreateForClass(ctx context.Context, assessment *models.Assessment) error
	FindByPublicID(ctx context.Context, publicID uuid.UUID) (*models.Class, error)
	FindAssessmentsByID(ctx context.Context, classID uint) ([]models.Assessment, error)
}

type classRepoImpl struct {
	DB *gorm.DB
}

func NewClassRepo(db *gorm.DB) ClassRepo {
	return &classRepoImpl{DB: db}
}

func (r *classRepoImpl) FindAssessmentsByID(ctx context.Context, classID uint) ([]models.Assessment, error) {
	var assessments []models.Assessment
	err := r.DB.WithContext(ctx).Find(&assessments, "class_id = ?", classID).Error
	return assessments, err
}

func (r *classRepoImpl) FindByPublicID(ctx context.Context, publicID uuid.UUID) (*models.Class, error) {
	var class models.Class
	err := r.DB.WithContext(ctx).First(&class, "public_id = ?", publicID).Error
	return &class, err
}

func (r *classRepoImpl) CreateForClass(ctx context.Context, assessment *models.Assessment) error {
	return r.DB.WithContext(ctx).Create(assessment).Error
}

func (r *classRepoImpl) Delete(ctx context.Context, classID uint, userID uint64) error {
	return r.DB.WithContext(ctx).
		Where("class_id = ? AND user_id = ?", classID, userID).
		Delete(&models.ClassUser{}).Error
}

func (r *classRepoImpl) FindByClassCode(ctx context.Context, classCode string) (*models.Class, error) {
	var class models.Class
	err := r.DB.WithContext(ctx).First(&class, "code = ?", classCode).Error
	return &class, err
}

func (r *classRepoImpl) IsUserJoined(ctx context.Context, classID uint, userID uint64) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.ClassUser{}).Where("class_id = ? AND user_id = ?", classID, userID).Count(&count).Error
	return count > 0, err
}

func (r *classRepoImpl) ExistByID(ctx context.Context, classID uint) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.Class{}).Where("id = ?", classID).Count(&count).Error
	return count > 0, err
}

func (r *classRepoImpl) ExistByClassCode(ctx context.Context, classCode string) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.Class{}).Where("code = ?", classCode).Count(&count).Error
	return count > 0, err
}

func (r *classRepoImpl) JoinClass(ctx context.Context, classUser *models.ClassUser) error {
	err := r.DB.WithContext(ctx).Create(classUser).Error
	return err
}

func (r *classRepoImpl) FindByUserID(ctx context.Context, userID uint64) ([]models.Class, error) {
	var classes []models.Class

	err := r.DB.WithContext(ctx).
		Model(&models.Class{}).
		Preload("CreatedBy").
		Preload("CreatedBy.Role").
		Joins("LEFT JOIN class_users cu ON cu.class_id = classes.id").
		Where(
			"classes.created_by_id = ? OR cu.user_id = ?",
			userID, userID,
		).
		Group("classes.id").
		Find(&classes).Error

	return classes, err
}

func (r *classRepoImpl) Create(ctx context.Context, class *models.Class) error {
	return r.DB.WithContext(ctx).Create(class).Error
}

func (r *classRepoImpl) CreateNewClass(ctx context.Context, class *models.Class) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(class).Error; err != nil {
			return err
		}

		err := tx.Exec("INSERT INTO class_users (class_id, user_id) VALUES (?, ?)", class.ID, class.CreatedById).Error
		if err != nil {
			return err
		}

		return nil
	})
}
