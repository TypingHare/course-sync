package fs

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
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

func EnsureDirExists(path string) error {
	exists, err := DirExists(path)
	if err != nil {
		return err
	}

	if !exists {
		return os.MkdirAll(path, 0o755)
	}

	return nil
}

// CollectFilesRecursively walks the directory tree rooted at dir and returns
// the paths of all files it finds.
//
// Entries whose base name appears in ignoredNames are skipped. If an ignored
// entry is a directory, its entire subtree is skipped. If it is a file, the
// file is ignored.
//
// The returned file paths are absolute or relative depending on the value of
// dir. If an error occurs while walking the directory tree, the walk stops and
// the error is returned.
func CollectFilesRecursively(
	dir string,
	ignoredFiles []string,
) ([]string, error) {
	var files []string

	err := filepath.WalkDir(
		dir,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if slices.Contains(ignoredFiles, d.Name()) {
				if d.IsDir() {
					return fs.SkipDir
				}
				return nil
			}

			if !d.IsDir() {
				files = append(files, path)
			}
			return nil
		},
	)

	return files, err
}
