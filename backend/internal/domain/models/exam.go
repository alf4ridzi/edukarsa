package models

import (
	"edukarsa-backend/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exam struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name string `gorm:"not null" json:"name"`

	StartAt time.Time `gorm:"not null" json:"start_at"`
	EndAt   time.Time `gorm:"not null" json:"end_at"`

	Duration int    `gorm:"not null" json:"duration"`
	Status   string `gorm:"not null;default:'draft'" json:"status"`

	ClassID uint   `json:"-"`
	Class   *Class `gorm:"foreignKey:ClassID" json:"class,omitempty"`

	Questions []ExamQuestion `gorm:"foreignKey:ExamID" json:"questions,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (e *Exam) BeforeCreate(tx *gorm.DB) (err error) {
	e.StartAt = e.StartAt.UTC()
	e.EndAt = e.EndAt.UTC()
	return
}

type ExamQuestion struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ExamID uuid.UUID `gorm:"type:uuid;index;not null" json:"-"`
	Exam   *Exam     `gorm:"foreignKey:ExamID" json:"exam,omitempty"`

	Question    string  `gorm:"type:text;not null" json:"question"`
	Explanation *string `json:"explanation"`

	AnswerID *uint       `json:"-"`
	Answer   *ExamOption `gorm:"-" json:"-"`

	Options []ExamOption `gorm:"foreignKey:ExamQuestionID" json:"options"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
}

type ExamSubmission struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	ExamID uuid.UUID `gorm:"type:uuid;index;not null"`
	UserID uint      `gorm:"index;not null"`

	StartAt     time.Time
	SubmittedAt *time.Time

	Status domain.SubmissionStatus `gorm:"not null;default:'ongoing'"` // submitted, ongoing, expired

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ExamScore struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	ExamID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_exam_user_score"`
	Exam   *Exam     `gorm:"foreignKey:ExamID"`

	UserID uint  `gorm:"not null;uniqueIndex:uq_exam_user_score"`
	User   *User `gorm:"foreignKey:UserID"`

	Correct int `gorm:"not null"`
	Wrong   int `gorm:"not null"`
	Score   int `gorm:"not null"`

	FinishedAt time.Time `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
