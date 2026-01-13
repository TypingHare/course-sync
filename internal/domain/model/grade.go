package model

import "time"

// Grade represents the grading information for a student's assignment.
type Grade struct {
	// AssignmentName is the name of the graded assignment.
	AssignmentName string `json:"assignment_name"`

	// SubmissionHash identifies the graded submission.
	SubmissionHash string `json:"submission_hash"`

	// Score is the grade received for the assignment.
	Score float64 `json:"score"`

	// Feedback contains instructor feedback for the assignment.
	Feedback string `json:"feedback"`

	// GradedAt is the time the assignment was graded.
	GradedAt time.Time `json:"graded_at"`
}
