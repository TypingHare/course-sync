package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const AppDataDirName = ".csync"

const GitRepositoryDirName = ".git"

const AssignmentDataFileName = "assignments.json"

const InstructorPrivateKeyFileName = "instructor"

const InstructorPublicKeyFileName = "instructor.pub"

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
		markerPath := filepath.Join(dir, GitRepositoryDirName)

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
		"git repository directory not found: %s (starting from %s)",
		GitRepositoryDirName,
		originalStart,
	)
}

func GetDataDir(projectDir string) string {
	return filepath.Join(projectDir, AppDataDirName)
}

func GetAssignmentDataFile(dataDir string) string {
	return filepath.Join(dataDir, AssignmentDataFileName)
}

func GetInstructorPrivateKeyFile(appDataDir string) string {
	return filepath.Join(appDataDir, InstructorPrivateKeyFileName)
}
