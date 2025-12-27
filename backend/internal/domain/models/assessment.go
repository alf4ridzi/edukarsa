package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assessment struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();uniqueIndex" json:"id"`
	Name       string    `json:"name" json:"name"`
	DeadlineAt time.Time `json:"deadline_at"`

	ClassID uint  `json:"-"`
	Class   Class `gorm:"foreignKey:ClassID;references:ID" json:"-"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type AssessmentCollection struct {
	ID       uint `gorm:"primaryKey"`
	File     string
	Feedback string
	Score    int

	AssessmentID uuid.UUID  `gorm:"type:uuid;not null;index"`
	Assessment   Assessment `gorm:"foreignKey:AssessmentID;references:ID"`

	UserID uint `gorm:"not null;index"`
	User   User `gorm:"foreignKey:UserID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
