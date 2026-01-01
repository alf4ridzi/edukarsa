package dto

type StudentAnswerRequest struct {
	OptionID uint `json:"option_id" binding:"required"`
}
