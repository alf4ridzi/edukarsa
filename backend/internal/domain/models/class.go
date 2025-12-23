package models

import (
	"edukarsa-backend/internal/utils"
	"errors"

	"gorm.io/gorm"
)

type Class struct {
	ID      uint   `gorm:"primaryKey"`
	ClassID string `gorm:"uniqueIndex:idx_class_id;size:10"`
	Name    string

	CreatedById uint
	CreatedBy   User `gorm:"foreignKey:CreatedById"`

	User []User `gorm:"many2many:class_users;"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) error {
	if c.ClassID != "" {
		return nil
	}

	for range 5 {
		id := utils.GenerateRandomString(8)

		var count int64
		tx.Model(&Class{}).Where("class_id = ?", id).Count(&count)

		if count == 0 {
			c.ClassID = id
			return nil
		}
	}

	return errors.New("failed to generate unique class id")
}
