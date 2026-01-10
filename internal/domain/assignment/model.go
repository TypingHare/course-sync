package assignment

import "time"

// Assignment represents the details of an academic assignment released to
// students.
type Assignment struct {
	// Short name or identifier for the assignment.
	Name string `json:"name"`

	// Full title of the assignment.
	Title string `json:"title"`

	// Detailed description of the assignment.
	Description string `json:"description"`

	// Time when the assignment is released.
	ReleasedAt time.Time `json:"releasedAt"`

	// Time when the assignment is due.
	DueAt time.Time `json:"dueAt"`

	// Maximum score achievable for the assignment.
	MaxScore float64 `json:"maxScore"`

	// Minimum score required to pass the assignment.
	PassingScore float64 `json:"passingScore"`
}
