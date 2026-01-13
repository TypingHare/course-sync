package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/service"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/filesystem"
	"github.com/TypingHare/course-sync/internal/support/io"
)

const AssignmentDataFileName = "assignments.json"

func GetAssignmentDataFile(dataDir string) string {
	return filepath.Join(dataDir, AssignmentDataFileName)
}

func GetAssignmentService(
	assignmentDataFile string,
) *service.AssignmentService {
	return service.NewAssignmentService(
		jsonstore.NewAssignmentRepo(assignmentDataFile),
	)
}

// GetUserAssignmentDir returns the path to the user assignment directory where
// the files for the specified assignment are stored.
func GetUserAssignmentDir(
	outputMode *io.OutputMode,
	srcDir string,
	assignmentName string,
) (string, error) {
	userWorkspaceDir, err := GetStudentWorkspaceDir(outputMode, srcDir)
	if err != nil {
		return "", fmt.Errorf("get user assignment dir: %w", err)
	}

	return filepath.Join(userWorkspaceDir, assignmentName), nil
}

func GetPrototypeAssignmentDir(
	srcDir string,
	assignmentName string,
) string {
	return filepath.Join(GetPrototypeWorkspaceDir(srcDir), assignmentName)
}

func PrepareAssignment(
	outputMode *io.OutputMode,
	projectDir string,
	assignmentName string,
	force bool,
) error {
	assignmentService := GetAssignmentService(
		GetAssignmentDataFile(
			GetDataDir(projectDir),
		),
	)

	// Find the assignment by name.
	assignment, err := assignmentService.GetAssignmentByName(assignmentName)
	if err != nil {
		return fmt.Errorf("get assignment: %w", err)
	}
	if assignment == nil {
		return fmt.Errorf("assignment %q not found", assignmentName)
	}

	// Get the source directory.
	srcDir := GetSrcDir(projectDir)

	// Check if the prototype assignment directory exists.
	prototypeAssignmentDir := GetPrototypeAssignmentDir(
		srcDir,
		assignmentName,
	)
	proptotypeAssignmentDirExists, err := filesystem.DirExists(
		prototypeAssignmentDir,
	)
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
	userAssignmentDir, err := GetUserAssignmentDir(
		outputMode,
		srcDir,
		assignmentName,
	)
	if err != nil {
		return fmt.Errorf("get user assignment directory: %w", err)
	}

	// If not forcing, check if the user assignment directory already exists.
	if !force {
		userAssignmentDirExists, err := filesystem.DirExists(userAssignmentDir)
		if err != nil {
			return fmt.Errorf("check user assignment directory: %w", err)
		}

		if userAssignmentDirExists {
			return fmt.Errorf(
				"user assignment directory %q already exists",
				userAssignmentDir,
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
