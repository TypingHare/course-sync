package assignment

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/workspace"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

// Assignments file name inside the application hidden directory.
const ASSIGNMENTS_FILE_NAME = "assignments.json"

// GetAssignments retrieves the list of assignments from the JSON file.
func GetAssignments(appCtx *app.Context) ([]Assignment, error) {
	assignments, err := jsonstore.ReadJSONFile[[]Assignment](
		filepath.Join(appCtx.AppDataDir, ASSIGNMENTS_FILE_NAME),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to read assignments: %w", err)
	}

	return assignments, nil
}

// FindAssignmentByName searches for an assignment by its name in the provided
// slice.
func FindAssignmentByName(assignments []Assignment, name string) *Assignment {
	for _, assignment := range assignments {
		if assignment.Name == name {
			return &assignment
		}
	}

	return nil
}

// GetUserAssignmentDir returns the path to the user assignment directory where
// the files for the specified assignment are stored.
func GetUserAssignmentDir(
	appCtx *app.Context,
	assignmentName string,
) (string, error) {
	userWorkspaceDir, err := workspace.GetUserWorkspaceDir(*appCtx)
	if err != nil {
		return "", fmt.Errorf("get user assignment dir: %w", err)
	}

	return filepath.Join(userWorkspaceDir, assignmentName), nil
}

// GetPrototypeAssignmentDir returns the path to the prototype directory for
// the specified assignment.
func GetPrototypeAssignmentDir(
	appCtx *app.Context,
	assignmentName string,
) string {
	return filepath.Join(
		workspace.GetPrototypeWorkspaceDir(*appCtx),
		assignmentName,
	)
}
