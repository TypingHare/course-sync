package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/TypingHare/course-sync/internal/infra/fs"
)

// CreateHashForDir returns a stable hash for the directory rooted at dir.
//
// The hash is computed by walking the directory tree, excluding entries whose
// base name appears in ignoredNames. For each regular file, the file’s relative
// path (from dir) and its contents are included in the hash. File paths are
// processed in sorted order to ensure deterministic output.
//
// The returned value is the last 12 hexadecimal characters of a SHA-256 digest.
func CreateHashForDir(dir string, ignoredNames []string) (string, error) {
	paths, err := fs.CollectFilesRecursively(dir, ignoredNames)
	if err != nil {
		return "", fmt.Errorf("collect submitted files: %w", err)
	}
	sort.Strings(paths)

	h := sha256.New()

	for _, p := range paths {
		if err := hashRegularFile(h, dir, p); err != nil {
			return "", err
		}
	}

	sum := hex.EncodeToString(h.Sum(nil))
	return sum[len(sum)-12:], nil
}

// hashRegularFile adds the relative path and contents of path to the hash.
//
// Only regular files are hashed; non-regular files (such as directories or
// symlinks) are ignored. The relative path is included to ensure that file
// renames or moves affect the resulting hash.
func hashRegularFile(h hash.Hash, root, path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("stat %q: %w", path, err)
	}
	if !info.Mode().IsRegular() {
		return nil
	}

	rel, err := filepath.Rel(root, path)
	if err != nil {
		return fmt.Errorf("rel %q: %w", path, err)
	}

	if err := writeLine(h, rel); err != nil {
		return fmt.Errorf("hash path %q: %w", rel, err)
	}
	if err := hashFile(h, path); err != nil {
		return fmt.Errorf("hash file %q: %w", path, err)
	}

	return nil
}

// writeLine writes s followed by a newline to the hash.
func writeLine(h hash.Hash, s string) error {
	if _, err := io.WriteString(h, s); err != nil {
		return err
	}
	_, err := io.WriteString(h, "\n")
	return err
}

// hashFile writes the contents of the file at path to the hash, followed by a
// newline.
func hashFile(h hash.Hash, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return err
	}
	_, err = io.WriteString(h, "\n")
	return err
}
