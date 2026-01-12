package assignment

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/TypingHare/course-sync/internal/domain/port"
	"github.com/TypingHare/course-sync/internal/domain/workspace"
)

// AssignmentsFileName is the name of the file where assignment records are
// stored.
const AssignmentsFileName = "assignments.json"

// CreateAssignment creates a new Assignment instance with the provided details.
func CreateAssignment(
	Name string,
	Title string,
	DueAt time.Time,
) *Assignment {
	return &Assignment{
		Name:         Name,
		Title:        Title,
		Description:  "",
		ReleasedAt:   time.Now().UTC(),
		DueAt:        DueAt,
		MaxScore:     100.0,
		PassingScore: 60.0,
	}
}

// FindAssignmentByName searches for an assignment by its name in the provided
// slice. It returns a pointer to the Assignment if found, or nil if not found.
func FindAssignmentByName(assignments []Assignment, name string) *Assignment {
	for i := range assignments {
		if assignments[i].Name == name {
			return &assignments[i]
		}
	}

	return nil
}

// GetUserAssignmentDir returns the path to the user assignment directory where
// the files for the specified assignment are stored.
func GetUserAssignmentDir(
	outputMode port.OutputMode,
	srcDir string,
	assignmentName string,
) (string, error) {
	userWorkspaceDir, err := workspace.GetUserWorkspaceDir(outputMode, srcDir)
	if err != nil {
		return "", fmt.Errorf("get user assignment dir: %w", err)
	}

	return filepath.Join(userWorkspaceDir, assignmentName), nil
}

// GetPrototypeAssignmentDir returns the path to the prototype directory for the
// specified assignment.
func GetPrototypeAssignmentDir(
	srcDir string,
	assignmentName string,
) string {
	return filepath.Join(
		workspace.GetPrototypeWorkspaceDir(srcDir),
		assignmentName,
	)
}
