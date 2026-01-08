package feature

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/execx"
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

// DirExists checks if a directory exists at the specified path.
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return info.IsDir(), nil
}

// MakeDirIfNotExists creates a directory at the specified path if it does not already exist.
func MakeDirIfNotExists(dirPath string) error {
	commandTask := execx.CommandTask{
		Command:        "mkdir",
		Args:           []string{"-p", dirPath},
		OngoingMessage: fmt.Sprintf("Creating directory <%s> if it does not exist...", dirPath),
		DoneMessage:    fmt.Sprintf("Directory <%s> created or already exists.", dirPath),
		ErrorMessage:   fmt.Sprintf("Failed to create directory <%s>.", dirPath),
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}

// DeleteDir deletes the directory at the specified path.
func DeleteDir(dirPath string) error {
	commandTask := execx.CommandTask{
		Command:        "rm",
		Args:           []string{"-rf", dirPath},
		OngoingMessage: fmt.Sprintf("Deleting directory <%s>...", dirPath),
		DoneMessage:    fmt.Sprintf("Directory <%s> deleted.", dirPath),
		ErrorMessage:   fmt.Sprintf("Failed to delete directory <%s>.", dirPath),
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}

// CopyDir copies a directory from sourceDirPath to destDirPath.
func CopyDir(sourceDirPath string, destDirPath string) error {
	commandTask := execx.CommandTask{
		Command: "cp",
		Args:    []string{"-r", sourceDirPath, destDirPath},
		OngoingMessage: fmt.Sprintf(
			"Copying directory from <%s> to <%s>...",
			sourceDirPath,
			destDirPath,
		),
		DoneMessage: fmt.Sprintf(
			"Directory copied from <%s> to <%s>.",
			sourceDirPath,
			destDirPath,
		),
		ErrorMessage: fmt.Sprintf(
			"Failed to copy directory from <%s> to <%s>.",
			sourceDirPath,
			destDirPath,
		),
		Quiet:        app.Quiet,
		PrintCommand: app.Verbose,
		PrintStdout:  app.Verbose,
		PrintStderr:  app.Verbose,
	}

	return commandTask.Start()
}
