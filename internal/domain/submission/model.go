package submission

import "time"

// Submission represents a student's submission for an assignment.
type Submission struct {
	// Hash is the unique identifier for the submission.
	Hash string `json:"hash"`

	// GitCommitHash is the git commit hash associated with the submission.
	GitCommitHash string `json:"git_hash"`

	// AssignmentName is the name of the associated assignment.
	AssignmentName string `json:"assignment_name"`

	// SubmittedAt is the time the submission was made.
	SubmittedAt time.Time `json:"submitted_at"`
}
