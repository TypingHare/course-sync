package workspace

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/domain/port"
	"github.com/TypingHare/course-sync/internal/domain/student"
	"github.com/TypingHare/course-sync/internal/infra/git"
)

// PROTOTYPE_WORKSPACE is the name of the prototype workspace directory.
const PROTOTYPE_WORKSPACE = "[prototype]"

// GetUserWorkspaceDir constructs the path to the user's workspace directory
// based on the git username and the project directory.
func GetUserWorkspaceDir(
	outputMode port.OutputMode,
	srcDir string,
) (string, error) {
	gitUsername, err := git.GetUsername(outputMode)
	if err != nil || gitUsername == "" {
		return "", fmt.Errorf("get git username: %w", err)
	}

	workspaceDirName := student.GetStudentDirName(gitUsername)
	return filepath.Join(srcDir, workspaceDirName), nil
}

// GetPrototypeWorkspaceDir constructs the path to the prototype workspace
// directory within the project directory.
func GetPrototypeWorkspaceDir(srcDir string) string {
	return filepath.Join(srcDir, PROTOTYPE_WORKSPACE)
}
