package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrProjectDirNotFound = errors.New("project directory not found")

// FindProjectDir walks upward from startDir (or cwd if empty) looking for the project root marker
// directory. It returns the directory containing the project root marker.
func FindProjectDir(startDir string) (string, error) {
	var err error
	origStart := startDir

	if startDir == "" {
		startDir, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("get current working directory: %w", err)
		}
	}

	// Resolve to an absolute, clean path
	dir, err := filepath.Abs(filepath.Clean(startDir))
	if err != nil {
		return "", fmt.Errorf("resolve start directory: %w", err)
	}

	// Best-effort: normalize symlinks
	if real, err := filepath.EvalSymlinks(dir); err == nil {
		dir = real
	}

	// If caller passed a file path, start from its parent directory
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

	if origStart == "" {
		origStart = startDir
	}

	return "", fmt.Errorf(
		"%w: %s (starting from %s)",
		ErrProjectDirNotFound,
		PROJECT_ROOT_MARKER_DIR_NAME,
		origStart,
	)
}
