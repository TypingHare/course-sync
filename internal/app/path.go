package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

// Error returned when the project directory cannot be found.
var ErrProjectDirNotFound = errors.New("project directory not found")

// Name of the directory that marks the project root. Since Course Sync is a CLI tool, usually the
// project root directory doesn't change in a lifecycle.
var projectDirCache string

// FindProjectDirPath walks upward from startDir (or cwd if empty) looking for the project root
// marker directory. It returns the path to the directory containing the project root marker.
func FindProjectDirPath(startDir string) (string, error) {
	var err error
	originalStart := startDir

	if startDir == "" {
		startDir, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("get current working directory: %w", err)
		}
	}

	// Resolve to an absolute, clean path.
	dir, err := filepath.Abs(filepath.Clean(startDir))
	if err != nil {
		return "", fmt.Errorf("resolve start directory: %w", err)
	}

	// Normalize symbolic links.
	if real, err := filepath.EvalSymlinks(dir); err == nil {
		dir = real
	}

	// If caller passed a file path, start from its parent directory.
	if fi, statErr := os.Stat(dir); statErr == nil && !fi.IsDir() {
		dir = filepath.Dir(dir)
	}

	for {
		markerPath := filepath.Join(dir, PROJECT_ROOT_MARKER_DIR_NAME)

		if fi, statErr := os.Stat(markerPath); statErr == nil && fi.IsDir() {
			return dir, nil
		} else if statErr != nil && !errors.Is(statErr, os.ErrNotExist) {
			return "", fmt.Errorf("stat project root marker %s: %w", markerPath, statErr)
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}

		dir = parent
	}

	if originalStart == "" {
		originalStart = startDir
	}

	return "", fmt.Errorf(
		"%w: %s (starting from %s)",
		ErrProjectDirNotFound,
		PROJECT_ROOT_MARKER_DIR_NAME,
		originalStart,
	)
}

func GetProjectDirPath() (string, error) {
	if projectDirCache != "" {
		return projectDirCache, nil
	}

	projectDir, err := FindProjectDirPath("")
	if err != nil {
		return "", err
	}

	projectDirCache = projectDir
	return projectDirCache, nil
}

func GetDocsDirPath() (string, error) {
	projectDir, err := GetProjectDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(projectDir, "docs"), nil
}

func GetSrcDirPath() (string, error) {
	projectDir, err := GetProjectDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(projectDir, "src"), nil
}

func GetPrototypeDirPath() (string, error) {
	projectDir, err := GetProjectDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(projectDir, "src", "[prototype]"), nil
}

func GetHomeDirPath(gitUsername string) (string, error) {
	projectDir, err := GetProjectDirPath()
	if err != nil {
		return "", err
	}

	dirName := strings.ReplaceAll(strcase.ToSnake(gitUsername), "_", "-")
	return filepath.Join(projectDir, "src", dirName), nil
}
