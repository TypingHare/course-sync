package filesystem_test

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/TypingHare/course-sync/internal/support/filesystem"
)

func TestDirExistsAndEnsureDirExists(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	missingDir := filepath.Join(tmp, "a", "b", "c")

	exists, err := filesystem.DirExists(missingDir)
	if err != nil {
		t.Fatalf("DirExists returned error for missing dir: %v", err)
	}
	if exists {
		t.Fatalf("DirExists = true for missing dir %q", missingDir)
	}

	if err := filesystem.EnsureDirExists(missingDir); err != nil {
		t.Fatalf("EnsureDirExists returned error: %v", err)
	}

	exists, err = filesystem.DirExists(missingDir)
	if err != nil {
		t.Fatalf("DirExists returned error for created dir: %v", err)
	}
	if !exists {
		t.Fatalf("DirExists = false for created dir %q", missingDir)
	}

	filePath := filepath.Join(tmp, "plain-file.txt")
	if err := os.WriteFile(filePath, []byte("x"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	exists, err = filesystem.DirExists(filePath)
	if err != nil {
		t.Fatalf("DirExists returned error for file path: %v", err)
	}
	if exists {
		t.Fatalf("DirExists = true for regular file %q", filePath)
	}
}

func TestCollectFilesRecursivelyHonorsIgnoredNames(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	mustWriteFile(t, filepath.Join(tmp, "keep.txt"), "keep")
	mustWriteFile(t, filepath.Join(tmp, "sub", "keep2.txt"), "keep2")
	mustWriteFile(t, filepath.Join(tmp, ".DS_Store"), "ignore")
	mustWriteFile(t, filepath.Join(tmp, "__pycache__", "ignored.pyc"), "ignore")

	files, err := filesystem.CollectFilesRecursively(
		tmp,
		[]string{".DS_Store", "__pycache__"},
	)
	if err != nil {
		t.Fatalf("CollectFilesRecursively returned error: %v", err)
	}

	rel := make([]string, 0, len(files))
	for _, f := range files {
		r, err := filepath.Rel(tmp, f)
		if err != nil {
			t.Fatalf("filepath.Rel failed: %v", err)
		}
		rel = append(rel, filepath.ToSlash(r))
	}
	sort.Strings(rel)

	want := []string{
		"keep.txt",
		"sub/keep2.txt",
	}

	if len(rel) != len(want) {
		t.Fatalf("got %d files (%v), want %d (%v)", len(rel), rel, len(want), want)
	}
	for i := range want {
		if rel[i] != want[i] {
			t.Fatalf("file[%d] = %q, want %q (all got: %v)", i, rel[i], want[i], rel)
		}
	}
}

func mustWriteFile(t *testing.T, path, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll(%q) failed: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile(%q) failed: %v", path, err)
	}
}
