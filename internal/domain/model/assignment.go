package model

import "time"

// Assignment represents the information of an assignment released to students.
type Assignment struct {
	// Name is a short identifier for the assignment.
	Name string `json:"name"`

	// Title is the full, human-readable assignment title.
	Title string `json:"title"`

	// Description provides detailed information about the assignment.
	Description string `json:"description"`

	// ReleasedAt is the time when the assignment becomes available.
	ReleasedAt time.Time `json:"released_at"`

	// DueAt is the deadline for submitting the assignment.
	DueAt time.Time `json:"due_at"`

	// MaxScore is the maximum achievable score.
	MaxScore float64 `json:"max_score"`

	// PassingScore is the minimum score required to pass.
	PassingScore float64 `json:"passing_score"`
}
