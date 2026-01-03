package domain

type SubmissionStatus string

const (
	SubmissionOngoing   SubmissionStatus = "ongoing"
	SubmissionExpired   SubmissionStatus = "expired"
	SubmissionSubmitted SubmissionStatus = "submitted"
)
