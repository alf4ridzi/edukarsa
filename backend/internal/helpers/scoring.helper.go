package helpers

import (
	"edukarsa-backend/internal/domain/dto"
	"edukarsa-backend/internal/domain/models"
)

func CalculateScore(questions []models.ExamQuestion, answers []models.ExamUserAnswer) dto.ExamScoreResult {
	answerMap := make(map[uint]uint)
	for _, answer := range answers {
		answerMap[answer.ExamQuestionID] = answer.AnswerID
	}

	var correct, wrong, unanswered int

	for _, question := range questions {
		answerID, ok := answerMap[question.ID]
		if !ok {
			unanswered++
			continue
		}

		if question.AnswerID != nil && answerID == *question.AnswerID {
			correct++
		} else {
			wrong++
		}
	}

	total := len(questions)
	score := 0
	if total > 0 {
		score = (correct * 100) / total
	}

	return dto.ExamScoreResult{
		Correct:    correct,
		Wrong:      wrong,
		UnAnswered: unanswered,
		Score:      score,
	}
}
