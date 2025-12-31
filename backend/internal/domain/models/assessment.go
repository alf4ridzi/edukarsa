package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assessment struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();uniqueIndex" json:"id"`
	Name       string    `json:"name"`
	DeadlineAt time.Time `json:"deadline_at"`

	ClassID uint  `json:"-"`
	Class   Class `gorm:"foreignKey:ClassID;references:ID" json:"-"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type AssessmentSubmission struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	FileUrl     string    `json:"file_url"`
	Feedback    *string   `json:"feedback"`
	Score       *int      `json:"score"`
	Description string    `json:"description"`

	AssessmentID uuid.UUID  `gorm:"type:uuid;not null;index" json:"-"`
	Assessment   Assessment `gorm:"foreignKey:AssessmentID;references:ID" json:"-"`

	UserID uint `gorm:"not null;index" json:"-"`
	User   User `gorm:"foreignKey:UserID;references:ID" json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
