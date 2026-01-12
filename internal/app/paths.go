package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Project root marker directory name. Folders containing this directory are
// considered project root directory (or simply project directory). The Git
// hidden directory is used as the project root marker because Course Sync is
// intended to be used in Git repositories.
const ProjectRootMarkerDirName = ".git"

// FindProjectDir walks upward from startDir looking for the project root marker
// directory. It returns the path to the directory containing the project root
// marker.
func FindProjectDir(startDir string) (string, error) {
	var err error
	originalStart := startDir

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
		markerPath := filepath.Join(dir, ProjectRootMarkerDirName)

		if fi, statErr := os.Stat(markerPath); statErr == nil && fi.IsDir() {
			return dir, nil
		} else if statErr != nil && !errors.Is(statErr, os.ErrNotExist) {
			return "", fmt.Errorf(
				"stat project root marker %s: %w",
				markerPath,
				statErr,
			)
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
		"project root marker not found: %s (starting from %s)",
		ProjectRootMarkerDirName,
		originalStart,
	)
}
