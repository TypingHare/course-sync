package doc_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func TestReleaseCmdRequiresPathForNewDoc(t *testing.T) {
	t.Parallel()

	projectDir := t.TempDir()
	ctx := testContext(projectDir)

	cmd := cli.Cmd(ctx)
	cmd.SetArgs([]string{"doc", "release", "guide"})

	err := cmd.Execute()
	if err == nil {
		t.Fatalf("release command succeeded without path for a new doc")
	}
	if !strings.Contains(err.Error(), "path is required") {
		t.Fatalf("error = %q, want missing-path message", err)
	}
}

func TestReleaseCmdCreatesNewDoc(t *testing.T) {
	t.Parallel()

	projectDir := t.TempDir()
	ctx := testContext(projectDir)

	cmd := cli.Cmd(ctx)
	cmd.SetArgs([]string{"doc", "release", "guide", "guide.md", "Guide", "v1"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("release command returned error: %v", err)
	}

	docs := readDocs(t, projectDir)
	if len(docs) != 1 {
		t.Fatalf("doc count = %d, want 1", len(docs))
	}

	got := docs[0]
	if got.Name != "guide" {
		t.Fatalf("doc name = %q, want %q", got.Name, "guide")
	}
	if got.Path != "guide.md" {
		t.Fatalf("doc path = %q, want %q", got.Path, "guide.md")
	}
	if got.Title != "Guide" {
		t.Fatalf("doc title = %q, want %q", got.Title, "Guide")
	}
	if got.Version != "v1" {
		t.Fatalf("doc version = %q, want %q", got.Version, "v1")
	}
	if got.IsDefault {
		t.Fatalf("doc is default = true, want false")
	}
	if got.ReleasedAt.IsZero() {
		t.Fatalf("released_at is zero")
	}
	if got.UpdatedAt.IsZero() {
		t.Fatalf("updated_at is zero")
	}
}

func TestReleaseCmdUpdatesExistingDocWithNameOnly(t *testing.T) {
	t.Parallel()

	projectDir := t.TempDir()
	existing := model.Doc{
		Name:       "guide",
		Path:       "guide.md",
		Title:      "Guide v1",
		Version:    "v1",
		ReleasedAt: time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC),
		UpdatedAt:  time.Date(2024, 1, 3, 3, 4, 5, 0, time.UTC),
		IsDefault:  true,
	}
	writeDocs(t, projectDir, []model.Doc{existing})

	ctx := testContext(projectDir)
	cmd := cli.Cmd(ctx)
	cmd.SetArgs([]string{"doc", "release", "guide"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("release command returned error: %v", err)
	}

	docs := readDocs(t, projectDir)
	if len(docs) != 1 {
		t.Fatalf("doc count = %d, want 1", len(docs))
	}

	got := docs[0]
	if got.Path != existing.Path {
		t.Fatalf("updated path = %q, want %q", got.Path, existing.Path)
	}
	if got.Title != existing.Title {
		t.Fatalf("updated title = %q, want %q", got.Title, existing.Title)
	}
	if got.Version != existing.Version {
		t.Fatalf("updated version = %q, want %q", got.Version, existing.Version)
	}
	if !got.ReleasedAt.Equal(existing.ReleasedAt) {
		t.Fatalf("released_at = %v, want %v", got.ReleasedAt, existing.ReleasedAt)
	}
	if !got.UpdatedAt.After(existing.UpdatedAt) {
		t.Fatalf("updated_at = %v, want after %v", got.UpdatedAt, existing.UpdatedAt)
	}
	if got.IsDefault != existing.IsDefault {
		t.Fatalf("is_default = %v, want %v", got.IsDefault, existing.IsDefault)
	}
}

func testContext(projectDir string) *app.Context {
	return &app.Context{
		OutputMode: io.NewOutputMode(false, true, true),
		WorkingDir: projectDir,
		ProjectDir: projectDir,
		Role:       model.RoleInstructor,
	}
}

func writeDocs(t *testing.T, projectDir string, docs []model.Doc) {
	t.Helper()

	dataFile := app.GetDocDataFile(app.GetDataDir(projectDir))
	if err := os.MkdirAll(filepath.Dir(dataFile), 0o755); err != nil {
		t.Fatalf("MkdirAll failed: %v", err)
	}

	payload, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent failed: %v", err)
	}
	if err := os.WriteFile(dataFile, payload, 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
}

func readDocs(t *testing.T, projectDir string) []model.Doc {
	t.Helper()

	dataFile := app.GetDocDataFile(app.GetDataDir(projectDir))
	payload, err := os.ReadFile(dataFile)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	var docs []model.Doc
	if err := json.Unmarshal(payload, &docs); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	return docs
}
