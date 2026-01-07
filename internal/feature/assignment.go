package feature

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
)

type Assignment struct {
	Name        string    `json:"name"`        // Short name or identifier for the assignment.
	Title       string    `json:"title"`       // Full title of the assignment.
	Description string    `json:"description"` // Detailed description of the assignment.
	ReleaseTime time.Time `json:"releaseTime"` // Time when the assignment is released.
	DueTime     time.Time `json:"dueTime"`     // Time when the assignment is due.
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
	appDirPath, err := app.GetAppDirPath()
	if err != nil {
		return nil, fmt.Errorf("get app dir: %w", err)
	}

	filePath := filepath.Join(appDirPath, app.ASSIGNMENTS_FILE_NAME)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", filePath, err)
	}

	var assignments []Assignment
	if err := json.Unmarshal(data, &assignments); err != nil {
		return nil, fmt.Errorf("parse %s: %w", filePath, err)
	}

	return assignments, nil
}

func SubmitAssignment(assignmentName string) error {
	userDirPath, err := GetUserDirPath()
	if err != nil {
		return fmt.Errorf("get user dir path: %w", err)
	}

	return GitAdd(filepath.Join(userDirPath, assignmentName))
}
