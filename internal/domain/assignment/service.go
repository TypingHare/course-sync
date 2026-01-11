package assignment

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/workspace"
	"github.com/TypingHare/course-sync/internal/infra/exec"
	"github.com/TypingHare/course-sync/internal/infra/fs"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

// Assignments file name inside the application data directory.
const ASSIGNMENTS_FILE_NAME = "assignments.json"

// GetAssignments retrieves the list of assignments from the assignments JSON
// file in the application data directory.
func GetAssignments(appCtx *app.Context) ([]Assignment, error) {
	assignments, err := jsonstore.ReadJSONFile[[]Assignment](
		filepath.Join(appCtx.AppDataDir, ASSIGNMENTS_FILE_NAME),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve assignments: %w", err)
	}

	return assignments, nil
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
	appCtx *app.Context,
	assignmentName string,
) (string, error) {
	userWorkspaceDir, err := workspace.GetUserWorkspaceDir(appCtx)
	if err != nil {
		return "", fmt.Errorf("get user assignment dir: %w", err)
	}

	return filepath.Join(userWorkspaceDir, assignmentName), nil
}

// GetPrototypeAssignmentDir returns the path to the prototype directory for the
// specified assignment.
func GetPrototypeAssignmentDir(
	appCtx *app.Context,
	assignmentName string,
) string {
	return filepath.Join(
		workspace.GetPrototypeWorkspaceDir(appCtx),
		assignmentName,
	)
}

// PrepareAssignment prepares the assignment environment by copying the
// prototype assignment to the user's assignment directory. If 'force' is true,
// it will overwrite any existing user assignment directory.
func PrepareAssignment(
	appCtx *app.Context,
	assignmentName string,
	force bool,
) error {
	// Find the assignment by name.
	assignments, err := GetAssignments(appCtx)
	if err != nil {
		return fmt.Errorf("prepare assignment: %w", err)
	}

	assignment := FindAssignmentByName(assignments, assignmentName)
	if assignment == nil {
		return fmt.Errorf("assignment %q not found", assignmentName)
	}

	// Check if the prototype assignment directory exists.
	prototypeAssignmentDir := GetPrototypeAssignmentDir(appCtx, assignmentName)
	proptotypeAssignmentDirExists, err := fs.DirExists(prototypeAssignmentDir)
	if err != nil {
		return fmt.Errorf("check prototype assignment directory: %w", err)
	}
	if !proptotypeAssignmentDirExists {
		return fmt.Errorf(
			"prototype assignment directory %q does not exist",
			prototypeAssignmentDir,
		)
	}

	// Check if the user assignment directory already exists.
	userAssignmentDir, err := GetUserAssignmentDir(appCtx, assignmentName)
	if err != nil {
		return fmt.Errorf("get user assignment directory: %w", err)
	}

	// If not forcing, check if the user assignment directory already exists.
	if !force {
		userAssignmentDirExists, err := fs.DirExists(userAssignmentDir)
		if err != nil {
			return fmt.Errorf("check user assignment directory: %w", err)
		}

		if userAssignmentDirExists {
			return fmt.Errorf(
				"user assignment directory %q already exists",
				assignmentName,
			)
		}
	}

	// Delete the existing user assignment directory if forcing.
	if force {
		err = exec.ShellDeleteDir(appCtx, userAssignmentDir)
		if err != nil {
			return fmt.Errorf(
				"delete user assignment directory: %w",
				err,
			)
		}
	}

	// Ensure the parent directory of the user assignment directory exists.
	err = exec.ShellEnsureDir(appCtx, filepath.Dir(userAssignmentDir))
	if err != nil {
		return fmt.Errorf(
			"ensure user assignment parent directory: %w",
			err,
		)
	}

	// Copy the prototype assignment to the user assignment directory.
	err = exec.ShellCopyDir(appCtx, prototypeAssignmentDir, userAssignmentDir)
	if err != nil {
		return fmt.Errorf("copy prototype assignment: %w", err)
	}

	return nil
}
