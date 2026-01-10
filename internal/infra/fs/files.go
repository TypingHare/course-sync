package fs

import (
	"errors"
	"os"
)

// FileExists checks if a file exists at the given path. It returns true if
// the file exists and is not a directory, false if it does not exist, and an
// error if there was an issue checking the file.
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
