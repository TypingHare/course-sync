package feature

import (
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/iancoleman/strcase"
)

// GetUserDirPath constructs the path to the user's directory within the source directory based on
// the Git user name. The directory name is derived by converting the Git user name to snake_case
// and replacing underscores with hyphens.
func GetUserDirPath() (string, error) {
	gitUsername, err := GetGitUserName()
	if err != nil || gitUsername == "" {
		return "", err
	}
	srcDirPath, err := app.GetSrcDirPath()
	if err != nil {
		return "", err
	}

	dirName := strings.ReplaceAll(strcase.ToSnake(gitUsername), "_", "-")
	userDirPath := filepath.Join(srcDirPath, dirName)

	return userDirPath, nil
}
