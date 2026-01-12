package assignmentservice

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/domain/assignment"
	"github.com/TypingHare/course-sync/internal/domain/port"
	"github.com/TypingHare/course-sync/internal/infra/exec"
	"github.com/TypingHare/course-sync/internal/infra/fs"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

// GetAssignments retrieves the list of assignments from the assignments JSON
// file in the application data directory.
func GetAssignments(appDataDir string) ([]assignment.Assignment, error) {
	assignments, err := jsonstore.ReadJSONFile[[]assignment.Assignment](
		filepath.Join(appDataDir, assignment.AssignmentsFileName),
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve assignments: %w", err)
	}

	return assignments, nil
}

// PrepareAssignment prepares the assignment environment by copying the
// prototype assignment to the user's assignment directory. If 'force' is true,
// it will overwrite any existing user assignment directory.
func PrepareAssignment(
	outputMode port.OutputMode,
	projectDir string,
	appDataDir string,
	srcDir string,
	assignmentName string,
	force bool,
) error {
	// Find the assignment by name.
	assignments, err := GetAssignments(appDataDir)
	if err != nil {
		return fmt.Errorf("prepare assignment: %w", err)
	}

	a := assignment.FindAssignmentByName(assignments, assignmentName)
	if a == nil {
		return fmt.Errorf("assignment %q not found", assignmentName)
	}

	// Check if the prototype assignment directory exists.
	prototypeAssignmentDir := assignment.GetPrototypeAssignmentDir(
		srcDir,
		assignmentName,
	)
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
	userAssignmentDir, err := assignment.GetUserAssignmentDir(
		outputMode,
		srcDir,
		assignmentName,
	)
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
		err = exec.ShellDeleteDir(outputMode, projectDir, userAssignmentDir)
		if err != nil {
			return fmt.Errorf(
				"delete user assignment directory: %w",
				err,
			)
		}
	}

	// Ensure the parent directory of the user assignment directory exists.
	err = exec.ShellEnsureDir(
		outputMode,
		projectDir,
		filepath.Dir(userAssignmentDir),
	)
	if err != nil {
		return fmt.Errorf(
			"ensure user assignment parent directory: %w",
			err,
		)
	}

	// Copy the prototype assignment to the user assignment directory.
	err = exec.ShellCopyDir(
		outputMode,
		projectDir,
		prototypeAssignmentDir,
		userAssignmentDir,
	)
	if err != nil {
		return fmt.Errorf("copy prototype assignment: %w", err)
	}

	return nil
}

// SaveAssignmentToFile appends the provided assignment to the assignments JSON
// file in the application data directory.
func SaveAssignmentToFile(
	outputMode port.OutputMode,
	appDataDir string,
	a *assignment.Assignment,
) error {
	assignments, err := GetAssignments(appDataDir)
	if err != nil {
		return fmt.Errorf("save assignment: %w", err)
	}

	assignments = append(assignments, *a)

	return jsonstore.WriteJSONFile(
		filepath.Join(appDataDir, assignment.AssignmentsFileName),
		assignments,
	)
}
