package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func GetUserWorkspaceDir(
	outputMode io.OutputMode,
	srcDir string,
) (string, error) {
	gitUsername, err := exec.GitGetUsername(outputMode)
	if err != nil || gitUsername == "" {
		return "", fmt.Errorf("get git username: %w", err)
	}

	studentName := gitUsername
	workspaceDirName := GetStudentDirName(studentName)

	return filepath.Join(srcDir, workspaceDirName), nil
}
