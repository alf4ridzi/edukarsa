package models

import (
	"github.com/google/uuid"
)

type TokenResponse struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type AssessmentSubmissionResponse struct {
	ID          uuid.UUID `json:"id"`
	FileURL     string    `json:"file_url"`
	Description string    `json:"description"`
}
