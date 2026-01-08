package feature

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
)

type Assignment struct {
	Name         string    `json:"name"`         // Short name or identifier for the assignment.
	Title        string    `json:"title"`        // Full title of the assignment.
	Description  string    `json:"description"`  // Detailed description of the assignment.
	ReleasedAt   time.Time `json:"releasedAt"`   // Time when the assignment is released.
	DueAt        time.Time `json:"dueAt"`        // Time when the assignment is due.
	MaxScore     float64   `json:"maxScore"`     // Maximum score achievable for the assignment.
	PassingScore float64   `json:"passingScore"` // Minimum score required to pass the assignment.
}

// GetAssignments loads all assignments from the application's assignments file.
//
// The assignments file is expected to live in the application directory returned by
// app.GetAppDirPath(). If the file does not exist (for example, on first run), an empty slice is
// returned and no error is reported.
//
// The returned slice contains fully unmarshaled Assignment values. Callers should validate the
// assignments if additional semantic checks are required.
func GetAssignments() ([]Assignment, error) {
	assignments, err := ReadJSONSlice[[]Assignment](app.ASSIGNMENTS_FILE_NAME)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}

// SubmitAssignment stages the specified assignment for commit in Git.
//
// The assignmentName parameter should correspond to the Name field of an Assignment.
func SubmitAssignment(assignmentName string) error {
	userDirPath, err := GetUserDirPath()
	if err != nil {
		return fmt.Errorf("get user dir path: %w", err)
	}

	return GitAdd(filepath.Join(userDirPath, assignmentName))
}
