package fs

import (
	"os"
)

// FileExists checks if a file exists at the given path. It returns true if
// the file exists and is not a directory, false if it does not exist, and an
// error if there was an issue checking the file.
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.Mode().IsRegular(), nil
}

// DirExists checks if a directory exists at the given path. It returns true if
// the directory exists, false if it does not exist, and an error if there was
// an issue checking the directory.
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.IsDir(), nil
}
