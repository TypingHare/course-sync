package feature

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
)

func ReadJSONSlice[T any](filename string) (T, error) {
	var zero T

	appDirPath, err := app.GetAppDirPath()
	if err != nil {
		return zero, fmt.Errorf("get app dir: %w", err)
	}

	filepath := filepath.Join(appDirPath, filename)

	data, err := os.ReadFile(filepath)
	if err != nil {
		return zero, fmt.Errorf("read %s: %w", filepath, err)
	}

	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return zero, fmt.Errorf("parse %s: %w", filepath, err)
	}

	return result, nil
}
