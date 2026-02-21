package hash_test

import (
	"os"
	"path/filepath"
	"testing"

	hashutil "github.com/TypingHare/course-sync/internal/support/hash"
)

func TestCreateHashForDirStableAndContentSensitive(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	writeHashTestFile(t, filepath.Join(tmp, "a.txt"), "alpha")
	writeHashTestFile(t, filepath.Join(tmp, "sub", "b.txt"), "beta")

	h1, err := hashutil.CreateHashForDir(tmp, nil)
	if err != nil {
		t.Fatalf("CreateHashForDir h1 returned error: %v", err)
	}
	h2, err := hashutil.CreateHashForDir(tmp, nil)
	if err != nil {
		t.Fatalf("CreateHashForDir h2 returned error: %v", err)
	}
	if h1 != h2 {
		t.Fatalf("hash should be stable; h1=%q h2=%q", h1, h2)
	}

	writeHashTestFile(t, filepath.Join(tmp, "sub", "b.txt"), "beta-updated")
	h3, err := hashutil.CreateHashForDir(tmp, nil)
	if err != nil {
		t.Fatalf("CreateHashForDir h3 returned error: %v", err)
	}
	if h3 == h1 {
		t.Fatalf("hash did not change after content change; before=%q after=%q", h1, h3)
	}
}

func TestCreateHashForDirIgnoresIgnoredNamesAndSymlink(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	baseFile := filepath.Join(tmp, "keep.txt")
	writeHashTestFile(t, baseFile, "stable")

	before, err := hashutil.CreateHashForDir(tmp, []string{"__pycache__", ".DS_Store"})
	if err != nil {
		t.Fatalf("CreateHashForDir(before) returned error: %v", err)
	}

	writeHashTestFile(t, filepath.Join(tmp, ".DS_Store"), "ignored")
	writeHashTestFile(t, filepath.Join(tmp, "__pycache__", "a.pyc"), "ignored")

	linkPath := filepath.Join(tmp, "keep-link.txt")
	if err := os.Symlink(baseFile, linkPath); err != nil {
		t.Skipf("Symlink not supported in environment: %v", err)
	}

	after, err := hashutil.CreateHashForDir(tmp, []string{"__pycache__", ".DS_Store"})
	if err != nil {
		t.Fatalf("CreateHashForDir(after) returned error: %v", err)
	}

	if before != after {
		t.Fatalf("ignored entries affected hash; before=%q after=%q", before, after)
	}
}

func writeHashTestFile(t *testing.T, path, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll(%q) failed: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile(%q) failed: %v", path, err)
	}
}
