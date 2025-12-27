package models

import "time"

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RegisterUser struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
}

type Login struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UpdateUserData struct {
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Username *string    `json:"username"`
	BirthDay *time.Time `json:"birthday"`
}

type CreateClassRequest struct {
	Name string `json:"name" binding:"required"`
}

type JoinClassRequest struct {
	ClassCode string `json:"code" binding:"required"`
}

type CreateAssessmentRequest struct {
	Name       string    `json:"name" binding:"required"`
	DeadlineAt time.Time `json:"deadline_at" binding:"required"`
}
