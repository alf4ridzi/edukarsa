package helpers

import (
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/models"
	"time"
)

func CanStudentGetExam(exam *models.Exam) error {
	if exam.Status != "published" && exam.Status != "ongoing" {
		return domain.ErrExamNotAccessible
	}

	return nil
}

func CanStudentStartExam(exam *models.Exam) error {
	now := time.Now().UTC()

	if exam.Status != "published" && exam.Status != "ongoing" {
		return domain.ErrExamNotAccessible
	}

	if now.Before(exam.StartAt) {
		return domain.ErrExamNotStarted
	}

	if !exam.EndAt.IsZero() && now.After(exam.EndAt) {
		return domain.ErrExamAlreadyFinished
	}

	return nil
}
