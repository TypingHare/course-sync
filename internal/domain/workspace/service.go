package workspace

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/git"
	"github.com/iancoleman/strcase"
)

// SRC_DIR_NAME is the name of the source directory within the project.
const SRC_DIR_NAME = "src"

// PROTOTYPE_WORKSPACE is the name of the prototype workspace directory.
const PROTOTYPE_WORKSPACE = "prototype"

// GetUserWorkspaceDir constructs the path to the user's workspace directory
// based on the git username and the project directory.
func GetUserWorkspaceDir(appCtx app.Context) (string, error) {
	gitUsername, err := git.GitGetUsername(&appCtx)
	if err != nil || gitUsername == "" {
		return "", fmt.Errorf("failed to get git username: %w", err)
	}

	workspaceDirName := strings.ReplaceAll(
		strcase.ToSnake(gitUsername),
		"_",
		"-",
	)
	workSpaceDir := filepath.Join(
		appCtx.ProjectDir,
		SRC_DIR_NAME,
		workspaceDirName,
	)

	return workSpaceDir, nil
}

// GetPrototypeWorkspaceDir constructs the path to the prototype workspace
// directory within the project directory.
func GetPrototypeWorkspaceDir(appCtx app.Context) string {
	return filepath.Join(
		appCtx.ProjectDir,
		SRC_DIR_NAME,
		PROTOTYPE_WORKSPACE,
	)
}
