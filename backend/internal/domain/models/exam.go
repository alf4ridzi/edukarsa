package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exam struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name     string    `gorm:"not null"`
	StartAt  time.Time `gorm:"not null"`
	Duration int       `gorm:"not null"`
	Status   string    `gorm:"not null;default:'draft'"`

	ClassID uint
	Class   Class `gorm:"foreignKey:ClassID"`

	Questions []ExamQuestion `gorm:"foreignKey:ExamID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ExamQuestion struct {
	ID uint `gorm:"primaryKey"`

	ExamID uuid.UUID `gorm:"type:uuid;index;not null"`
	Exam   *Exam     `gorm:"foreignKey:ExamID"`

	Question    string `gorm:"type:text;not null"`
	Explanation *string

	AnswerID *uint       `json:"-"`
	Answer   *ExamOption `gorm:"-"`

	Options []ExamOption `gorm:"foreignKey:ExamQuestionID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ExamOption struct {
	ID uint `gorm:"primaryKey"`

	ExamQuestionID uint          `gorm:"index;not null"`
	ExamQuestion   *ExamQuestion `gorm:"foreignKey:ExamQuestionID"`

	Option string `gorm:"type:text;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ExamUserAnswer struct {
	ID uint `gorm:"primaryKey"`

	ExamID uuid.UUID `gorm:"type:uuid;index;not null"`
	Exam   Exam      `gorm:"foreignKey:ExamID"`

	UserID uint `gorm:"index;not null"`
	User   User `gorm:"foreignKey:UserID"`

	ExamQuestionID uint         `gorm:"index;not null"`
	ExamQuestion   ExamQuestion `gorm:"foreignKey:ExamQuestionID"`

	AnswerID uint
	Answer   ExamOption `gorm:"foreignKey:AnswerID"`

	CreatedAt time.Time
	UpdatedAt time.Time

	_ struct{} `gorm:"uniqueIndex:uq_exam_user_question,priority:1"`
}

type ExamScore struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	ExamID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_exam_user_score"`
	Exam   Exam      `gorm:"foreignKey:ExamID"`

	UserID uint `gorm:"not null;uniqueIndex:uq_exam_user_score"`
	User   User `gorm:"foreignKey:UserID"`

	Correct int `gorm:"not null"`
	Wrong   int `gorm:"not null"`
	Score   int `gorm:"not null"`

	FinishedAt time.Time `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
