package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"size:50;uniqueIndex:idx_role;not null" json:"name"`
}
