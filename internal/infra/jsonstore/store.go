package jsonstore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ReadJSONFile reads the file at filePath and decodes exactly one JSON values
// into a value of type T.
//
// The function is intentionally strict:
//
//   - Unknown JSON fields are rejected.
//   - Trailing JSON or non-whitespace data after the first value is rejected.
//
// If the file does not exist, cannot be read, contains invalid JSON, contains
// unknown fields, or contains trailing data, an error is returned and the zero
// value of T is returned alongside the error.
func ReadJSONFile[T any](filePath string) (T, error) {
	var zero T

	data, err := os.ReadFile(filePath)
	if err != nil {
		return zero, fmt.Errorf("read %s: %w", filePath, err)
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	var value T
	if err := dec.Decode(&value); err != nil {
		return zero, fmt.Errorf("decode %s: %w", filePath, err)
	}

	if err := dec.Decode(new(any)); err != io.EOF {
		return zero, fmt.Errorf("decode %s: trailing data", filePath)
	}

	return value, nil
}

// WriteJSONFile marshals value as pretty-printed JSON and writes it to
// filePath.
//
// The write is performed atomically by writing the JSON data to a temporary
// file in the same directory and then renaming it over the target path. This
// ensures that the target file is either fully written or not modified at all.
//
// Parent directories are created if they do not already exist. If an error
// occurs at any stage, the original file (if any) is left unchanged.
//
// The function returns an error if the directory cannot be created, the value
// cannot be marshaled, the temporary file cannot be written or synced, or the
// final rename fails.
func WriteJSONFile[T any](filePath string, value T) error {
	// Ensure parent directory exists.
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create dir for %s: %w", filePath, err)
	}

	// Marshal with indentation for readability.
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", filePath, err)
	}

	// Create a temporary file in the same directory for atomic replace.
	tmp, err := os.CreateTemp(dir, ".json-*.tmp")
	if err != nil {
		return fmt.Errorf("create temp file for %s: %w", filePath, err)
	}
	tmpName := tmp.Name()

	// Ensure temporary file is removed on failure.
	defer func() {
		_ = os.Remove(tmpName)
	}()

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("write temp file %s: %w", tmpName, err)
	}

	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("sync temp file %s: %w", tmpName, err)
	}

	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close temp file %s: %w", tmpName, err)
	}

	// Atomically replace the target file.
	if err := os.Rename(tmpName, filePath); err != nil {
		return fmt.Errorf("replace %s: %w", filePath, err)
	}

	return nil
}
