package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/io"
)

const PrototypeWorkspaceName = "[prototype]"

// GetWorkspaceDir returns the workspace directory path for the given source
// dir and workspace name.
func GetWorkspaceDir(srcDir string, workspaceName string) string {
	return filepath.Join(srcDir, workspaceName)
}

// GetStudentWorkspaceDir returns the current student's workspace directory
// using the git username.
func GetStudentWorkspaceDir(
	outputMode *io.OutputMode,
	srcDir string,
) (string, error) {
	gitUsername, err := exec.GitGetUsername(outputMode)
	if err != nil || gitUsername == "" {
		return "", fmt.Errorf("get git username: %w", err)
	}

	studentDirName := GetStudentDirName(gitUsername)
	return GetWorkspaceDir(srcDir, studentDirName), nil
}

// GetPrototypeWorkspaceDir returns the prototype workspace directory path.
func GetPrototypeWorkspaceDir(srcDir string) string {
	return GetWorkspaceDir(srcDir, PrototypeWorkspaceName)
}
