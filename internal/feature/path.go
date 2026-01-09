package feature

import (
	"errors"
	"fmt"
	"io/fs"
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
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return !info.IsDir(), nil
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
	relativePath, err := app.GetRelativePath(dirPath)
	if err != nil {
		return errors.New("failed to get relative path: " + err.Error())
	}

	commandTask := execx.CommandTask{
		Command: "mkdir",
		Args:    []string{"-p", dirPath},
		OngoingMessage: fmt.Sprintf(
			"Creating directory <%s> if it does not exist...",
			relativePath,
		),
		DoneMessage:  fmt.Sprintf("Directory <%s> created or already exists.", relativePath),
		ErrorMessage: fmt.Sprintf("Failed to create directory <%s>.", relativePath),
		Quiet:        app.Quiet,
		PrintCommand: app.Verbose,
		PrintStdout:  app.Verbose,
		PrintStderr:  app.Verbose,
	}

	return commandTask.Start()
}

// DeleteDir deletes the directory at the specified path.
func DeleteDir(dirPath string) error {
	relativePath, err := app.GetRelativePath(dirPath)
	if err != nil {
		return errors.New("failed to get relative path: " + err.Error())
	}

	commandTask := execx.CommandTask{
		Command:        "rm",
		Args:           []string{"-rf", dirPath},
		OngoingMessage: fmt.Sprintf("Deleting directory <%s>...", relativePath),
		DoneMessage:    fmt.Sprintf("Directory <%s> deleted.", relativePath),
		ErrorMessage:   fmt.Sprintf("Failed to delete directory <%s>.", relativePath),
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}

// CopyFile copies a file to a destination directory using the "cp" command.
func CopyFile(sourceFiPath string, destDirPath string, force bool) error {
	relativeSourcePath, err := app.GetRelativePath(sourceFiPath)
	if err != nil {
		return errors.New("failed to get relative source directory path: " + err.Error())
	}

	relativeDestDirPath, err := app.GetRelativePath(destDirPath)
	if err != nil {
		return errors.New("failed to get relative destination directory path: " + err.Error())
	}

	args := []string{relativeSourcePath, destDirPath}
	if force {
		args = append([]string{"-f"}, args...)
	}

	commandTask := execx.CommandTask{
		Command: "cp",
		Args:    args,
		OngoingMessage: fmt.Sprintf(
			"Copying file <%s> to <%s>...",
			relativeSourcePath,
			relativeDestDirPath,
		),
		DoneMessage: fmt.Sprintf(
			"File copied from <%s> to <%s>.",
			relativeSourcePath,
			relativeDestDirPath,
		),
		ErrorMessage: fmt.Sprintf(
			"Failed to copy file <%s> to <%s>.",
			relativeSourcePath,
			relativeDestDirPath,
		),
		Quiet:        app.Quiet,
		PrintCommand: app.Verbose,
		PrintStdout:  app.Verbose,
		PrintStderr:  app.Verbose,
	}

	return commandTask.Start()
}

// CopyDir copies a directory from sourceDirPath to destDirPath.
func CopyDir(sourceDirPath string, destDirPath string) error {
	relativeSourceDirPath, err := app.GetRelativePath(sourceDirPath)
	if err != nil {
		return errors.New("failed to get relative source path: " + err.Error())
	}

	relativeDestDirPath, err := app.GetRelativePath(destDirPath)
	if err != nil {
		return errors.New("failed to get relative destination path: " + err.Error())
	}

	commandTask := execx.CommandTask{
		Command: "cp",
		Args:    []string{"-r", sourceDirPath, destDirPath},
		OngoingMessage: fmt.Sprintf(
			"Copying directory from <%s> to <%s>...",
			relativeSourceDirPath,
			relativeDestDirPath,
		),
		DoneMessage: fmt.Sprintf(
			"Directory copied from <%s> to <%s>.",
			relativeSourceDirPath,
			relativeDestDirPath,
		),
		ErrorMessage: fmt.Sprintf(
			"Failed to copy directory from <%s> to <%s>.",
			relativeSourceDirPath,
			relativeDestDirPath,
		),
		Quiet:        app.Quiet,
		PrintCommand: app.Verbose,
		PrintStdout:  app.Verbose,
		PrintStderr:  app.Verbose,
	}

	return commandTask.Start()
}

// CollectFiles collects all files in the specified directory and its subdirectories, excluding
// directories named "__pycache__" and files named ".DS_Store". It returns the absolute paths of the
// collected files.
func CollectFiles(dirPath string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dirPath, func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dirEntry.IsDir() && dirEntry.Name() == "__pycache__" {
			return fs.SkipDir
		}

		if !dirEntry.IsDir() && dirEntry.Name() != ".DS_Store" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
