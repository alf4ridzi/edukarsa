package models

import "time"

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
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

type UpdateUserData struct {
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Username *string    `json:"username"`
	BirthDay *time.Time `json:"birthday"`
}

type CreateClassRequest struct {
	Name string `json:"name"`
}

type JoinClassRequest struct {
	ClassID string `json:"class_id"`
}
