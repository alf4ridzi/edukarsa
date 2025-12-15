package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID   uint
	Role     Role `json:"role"`
	Name     string
	Email    string `gorm:"uniqueIndex:idx_email;size:50"`
	Username string `gorm:"uniqueIndex:idx_username;size:100"`
	Password string
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
