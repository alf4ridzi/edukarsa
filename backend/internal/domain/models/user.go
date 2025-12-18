package models

import (
	"edukarsa-backend/internal/utils"
	"html"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID   uint   `json:"-"`
	Role     Role   `json:"role"`
	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex:idx_email;size:50" json:"email"`
	Username string `gorm:"uniqueIndex:idx_username;size:100" json:"username"`
	Password string `json:"-"`
}

type RegisterUser struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type Login struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
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
