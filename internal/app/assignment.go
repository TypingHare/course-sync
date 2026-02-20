package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/service"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/filesystem"
	"github.com/TypingHare/course-sync/internal/support/io"
)

const AssignmentDataFileName = "assignments.json"

// GetAssignmentDataFile returns the assignments data file path under the given
// data directory.
func GetAssignmentDataFile(dataDir string) string {
	return filepath.Join(dataDir, AssignmentDataFileName)
}

// GetAssignmentService constructs an AssignmentService backed by the given
// data file.
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

// GetPrototypeAssignmentDir returns the prototype assignment directory path.
func GetPrototypeAssignmentDir(
	srcDir string,
	assignmentName string,
) string {
	return filepath.Join(GetPrototypeWorkspaceDir(srcDir), assignmentName)
}

// PrepareAssignment copies the prototype assignment into the user's workspace.
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
	prototypeAssignmentDirExists, err := filesystem.DirExists(
		prototypeAssignmentDir,
	)
	if err != nil {
		return fmt.Errorf("check prototype assignment directory: %w", err)
	}
	if !prototypeAssignmentDirExists {
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

// Assign registers a new assignment and distributes related files to student
// repos.
func Assign(
	outputMode *io.OutputMode,
	projectDir string,
	newAssignment *model.Assignment,
) error {
	assignmentDataFile := GetAssignmentDataFile(GetDataDir(projectDir))
	assignmentService := GetAssignmentService(assignmentDataFile)

	err := assignmentService.AddAssignment(newAssignment)
	if err != nil {
		return fmt.Errorf("add assignment: %w", err)
	}

	// Check if the prototype assignment directory exists.
	prototypeAssignmentDir := GetPrototypeAssignmentDir(
		GetSrcDir(projectDir),
		newAssignment.Name,
	)
	prototypeAssignmentDirExists, err := filesystem.DirExists(
		prototypeAssignmentDir,
	)
	if err != nil {
		return fmt.Errorf("check prototype assignment directory: %w", err)
	} else if !prototypeAssignmentDirExists {
		return fmt.Errorf(
			"prototype assignment directory %q does not exist",
			prototypeAssignmentDir,
		)
	}

	// Get all students.
	studentService := GetStudentService(
		GetStudentDataFile(GetDataDir(projectDir)),
	)
	students, err := studentService.GetAllStudents()
	if err != nil {
		return fmt.Errorf("get students: %w", err)
	}

	// Distribute the assignment data file to student repositories.
	assignmentDataFileRelPath, err := filepath.Rel(
		projectDir,
		assignmentDataFile,
	)
	if err != nil {
		return fmt.Errorf(
			"get relative path of assignment data file: %w",
			err,
		)
	}

	err = DistributeFileToStudentRepos(
		outputMode,
		students,
		projectDir,
		assignmentDataFileRelPath,
	)
	if err != nil {
		return fmt.Errorf(
			"distribute assignment to student repositories: %w",
			err,
		)
	}

	// Distribute the prototype assignment directory to student
	// repositories.
	prototypeAssignmentDirRelPath, err := filepath.Rel(
		projectDir,
		prototypeAssignmentDir,
	)
	if err != nil {
		return fmt.Errorf(
			"get relative path of prototype assignment directory: %w",
			err,
		)
	}

	err = DistributeDirToStudentRepos(
		outputMode,
		students,
		projectDir,
		prototypeAssignmentDirRelPath,
	)
	if err != nil {
		return fmt.Errorf(
			"distribute assignment directory to student repositories: %w",
			err,
		)
	}

	return nil
}
