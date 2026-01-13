package jsonstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type JsonStore[T any] struct {
	filePath string
}

func NewJsonStore[T any](filePath string) *JsonStore[T] {
	return &JsonStore[T]{filePath: filePath}
}

func (s *JsonStore[T]) Read() (T, error) {
	var zero T

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return zero, nil
		}
		return zero, fmt.Errorf("read %s: %w", s.filePath, err)
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	var value T
	if err := dec.Decode(&value); err != nil {
		return zero, fmt.Errorf("decode %s: %w", s.filePath, err)
	}

	if err := dec.Decode(new(any)); err != io.EOF {
		return zero, fmt.Errorf("decode %s: trailing data", s.filePath)
	}

	return value, nil
}

func (s *JsonStore[T]) Write(value T) error {
	// Ensure parent directory exists.
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create dir for %s: %w", s.filePath, err)
	}

	// Marshal with indentation for readability.
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", s.filePath, err)
	}

	// Create a temporary file in the same directory for atomic replace.
	tmp, err := os.CreateTemp(dir, ".json-*.tmp")
	if err != nil {
		return fmt.Errorf("create temp file for %s: %w", s.filePath, err)
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
	if err := os.Rename(tmpName, s.filePath); err != nil {
		return fmt.Errorf("replace %s: %w", s.filePath, err)
	}

	return nil
}
