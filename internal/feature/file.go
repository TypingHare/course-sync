package feature

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
)

// ReadJSONSlice reads a JSON file and unmarshals its content into a slice of type T. It returns the
// slice and any error encountered during the process.
func ReadJSONSlice[T any](filename string) (T, error) {
	var zero T

	appDirPath, err := app.GetAppDirPath()
	if err != nil {
		return zero, fmt.Errorf("get app dir: %w", err)
	}

	_filepath := filepath.Join(appDirPath, filename)

	data, err := os.ReadFile(_filepath)
	if err != nil {
		return zero, fmt.Errorf("read %s: %w", _filepath, err)
	}

	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return zero, fmt.Errorf("parse %s: %w", _filepath, err)
	}

	return result, nil
}

// WriteJSONSlice marshals the provided slice of type T into JSON format and writes it to a file.
// It ensures the write operation is atomic and returns any error encountered during the process.
func WriteJSONSlice[T any](filename string, data T) error {
	appDirPath, err := app.GetAppDirPath()
	if err != nil {
		return fmt.Errorf("get app dir: %w", err)
	}

	_filepath := filepath.Join(appDirPath, filename)

	// Ensure parent directory exists.
	if err := os.MkdirAll(filepath.Dir(_filepath), 0o755); err != nil {
		return fmt.Errorf("create dir for %s: %w", _filepath, err)
	}

	// Marshal with indentation for readability.
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", _filepath, err)
	}

	// Write atomically: write to temp file, then rename.
	tmpFile := _filepath + ".tmp"

	if err := os.WriteFile(tmpFile, jsonData, 0o644); err != nil {
		return fmt.Errorf("write temp file %s: %w", tmpFile, err)
	}

	if err := os.Rename(tmpFile, _filepath); err != nil {
		return fmt.Errorf("replace %s: %w", _filepath, err)
	}

	return nil
}
