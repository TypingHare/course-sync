package config

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
)

// GetConfigFilePath returns the absolute path to the configuration file.
func GetConfigFilePath() (string, error) {
	dir, err := app.GetProjectDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, app.CONFIG_FILE_NAME), nil
}
