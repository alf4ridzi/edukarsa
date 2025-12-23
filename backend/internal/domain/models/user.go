package models

import (
	"edukarsa-backend/internal/utils"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint       `gorm:"primarykey" json:"id"`
	RoleID   uint       `json:"-"`
	Role     Role       `json:"role"`
	Name     string     `json:"name"`
	Email    string     `gorm:"uniqueIndex:idx_email;size:50" json:"email"`
	Username string     `gorm:"uniqueIndex:idx_username;size:100" json:"username"`
	Password string     `json:"-"`
	BirthDay *time.Time `json:"birthday"`

	CreatedClass []Class `gorm:"foreignKey:CreatedByID"`
	Class        []Class `gorm:"many2many:class_users;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.RoleID == 0 {
		u.RoleID = 2
	}

	hashPass, err := utils.HashPasswordBcrypt(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashPass
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
