package models

import (
	"edukarsa-backend/internal/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"uniqueIndex:idx_class_code;size:10" json:"code"`
	Name string `json:"name"`

	CreatedById uint `json:"-"`
	CreatedBy   User `gorm:"foreignKey:CreatedById" json:"created_by"`

	Users []User `gorm:"many2many:class_users;" json:"user"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type ClassUser struct {
	UserID    uint `gorm:"primaryKey"`
	ClassID   uint `gorm:"primaryKey"`
	CreatedAt time.Time
}

func (c *Class) BeforeCreate(tx *gorm.DB) error {
	if c.Code != "" {
		return nil
	}

	for range 5 {
		code := utils.GenerateRandomString(8)

		var count int64
		tx.Model(&Class{}).Where("code = ?", code).Count(&count)

		if count == 0 {
			c.Code = code
			return nil
		}
	}

	return errors.New("failed to generate unique class id")
}
