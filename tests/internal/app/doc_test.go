package app_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func TestReleaseDocReturnsDocDistributionError(t *testing.T) {
	t.Parallel()

	projectDir := t.TempDir()
	dataDir := app.GetDataDir(projectDir)
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		t.Fatalf("MkdirAll data dir failed: %v", err)
	}

	students := []model.Student{{ID: 1, Name: "Alice"}}
	studentsData, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent students failed: %v", err)
	}
	if err := os.WriteFile(app.GetStudentDataFile(dataDir), studentsData, 0o644); err != nil {
		t.Fatalf("WriteFile students failed: %v", err)
	}

	err = app.ReleaseDoc(
		io.NewOutputMode(false, true, true),
		projectDir,
		&model.Doc{
			Name:       "guide",
			Path:       "missing.md",
			Title:      "Guide",
			Version:    "v1",
			ReleasedAt: time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
		},
	)
	if err == nil {
		t.Fatalf("ReleaseDoc succeeded for missing doc file; want error")
	}
	if !strings.Contains(err.Error(), "distribute doc file") {
		t.Fatalf("error = %q, want distribution context", err)
	}

	docDataFileInStudentRepo := filepath.Join(
		app.GetStudentRepoDir(projectDir, "Alice"),
		app.AppDataDirName,
		app.DocDataFileName,
	)
	if _, statErr := os.Stat(docDataFileInStudentRepo); statErr != nil {
		t.Fatalf("expected doc data file in student repo, stat failed: %v", statErr)
	}
}
