package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/io"
)

const PrototypeWorkspaceName = "[prototype]"

func GetWorkspaceDir(srcDir string, workspaceName string) string {
	return filepath.Join(srcDir, workspaceName)
}

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

func GetPrototypeWorkspaceDir(srcDir string) string {
	return GetWorkspaceDir(srcDir, PrototypeWorkspaceName)
}
