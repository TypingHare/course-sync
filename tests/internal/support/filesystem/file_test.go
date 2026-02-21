package filesystem_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/TypingHare/course-sync/internal/support/filesystem"
)

func TestFileExists(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	missing := filepath.Join(tmp, "missing.txt")

	exists, err := filesystem.FileExists(missing)
	if err != nil {
		t.Fatalf("FileExists returned error for missing file: %v", err)
	}
	if exists {
		t.Fatalf("FileExists = true for missing file %q", missing)
	}

	filePath := filepath.Join(tmp, "present.txt")
	if err := os.WriteFile(filePath, []byte("hello"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	exists, err = filesystem.FileExists(filePath)
	if err != nil {
		t.Fatalf("FileExists returned error for regular file: %v", err)
	}
	if !exists {
		t.Fatalf("FileExists = false for regular file %q", filePath)
	}

	exists, err = filesystem.FileExists(tmp)
	if err != nil {
		t.Fatalf("FileExists returned error for directory path: %v", err)
	}
	if exists {
		t.Fatalf("FileExists = true for directory path %q", tmp)
	}
}

func TestCopyFile(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	src := filepath.Join(tmp, "src.txt")
	dest := filepath.Join(tmp, "nested", "dir", "dest.txt")

	want := "payload"
	if err := os.WriteFile(src, []byte(want), 0o644); err != nil {
		t.Fatalf("WriteFile(src) failed: %v", err)
	}

	if err := filesystem.CopyFile(src, dest); err != nil {
		t.Fatalf("CopyFile returned error: %v", err)
	}

	gotBytes, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("ReadFile(dest) failed: %v", err)
	}
	if string(gotBytes) != want {
		t.Fatalf("copied content = %q, want %q", string(gotBytes), want)
	}
}

func TestRelOrOriginal(t *testing.T) {
	t.Parallel()

	base := "/tmp/base"
	path := "/tmp/base/a/b.txt"

	got := filesystem.RelOrOriginal(base, path)
	if got != filepath.Join("a", "b.txt") {
		t.Fatalf("RelOrOriginal(%q, %q) = %q, want %q", base, path, got, filepath.Join("a", "b.txt"))
	}
}
