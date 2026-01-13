package filesystem

import (
	"io"
	"os"
	"path/filepath"
)

// FileExists reports whether path exists and is a regular file.
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

// CopyFile copies the file at srcPath to destPath, creating parent dirs as
// needed.
func CopyFile(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destDir := filepath.Dir(destPath)
	if err := EnsureDirExists(destDir); err != nil {
		return err
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
