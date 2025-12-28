package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"size:50;uniqueIndex:idx_role;not null" json:"name"`

	// Permission []Permission `gorm:"many2many:role_permissions;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
