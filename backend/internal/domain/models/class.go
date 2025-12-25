package models

import (
	"edukarsa-backend/internal/utils"
	"errors"

	"gorm.io/gorm"
)

type Class struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"uniqueIndex:idx_class_code;size:10" json:"code"`
	Name string `json:"name"`

	CreatedById uint `json:"-"`
	CreatedBy   User `gorm:"foreignKey:CreatedById" json:"created_by"`

	User []User `gorm:"many2many:class_users;" json:"user"`
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
