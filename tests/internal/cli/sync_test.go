package cli_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func TestSyncInstructorCommitsAndPushesStudentRepos(t *testing.T) {
	projectDir := t.TempDir()
	dataDir := app.GetDataDir(projectDir)
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		t.Fatalf("MkdirAll data dir failed: %v", err)
	}

	students := []model.Student{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	studentsPayload, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent students failed: %v", err)
	}
	if err := os.WriteFile(app.GetStudentDataFile(dataDir), studentsPayload, 0o644); err != nil {
		t.Fatalf("WriteFile students failed: %v", err)
	}

	type repoFixture struct {
		repoDir    string
		remoteDir  string
		branchName string
	}
	fixtures := make(map[string]repoFixture, len(students))

	for i := range students {
		repoDir := app.GetStudentRepoDir(projectDir, students[i].Name)
		remoteDir := filepath.Join(projectDir, "remotes", fmt.Sprintf("%s.git", app.GetStudentDirName(students[i].Name)))

		branchName := setupRepoWithRemote(t, repoDir, remoteDir)
		fixtures[students[i].Name] = repoFixture{
			repoDir:    repoDir,
			remoteDir:  remoteDir,
			branchName: branchName,
		}
	}

	// Create changes in each repo to ensure sync stages, commits, and pushes.
	appendFile(t, filepath.Join(fixtures["Alice"].repoDir, "notes.txt"), "\nchange from alice\n")
	appendFile(t, filepath.Join(fixtures["Bob"].repoDir, "notes.txt"), "\nchange from bob\n")

	ctx := &app.Context{
		OutputMode: io.NewOutputMode(false, true, true),
		ProjectDir: projectDir,
		WorkingDir: projectDir,
		Role:       model.RoleInstructor,
	}

	cmd := cli.Cmd(ctx)
	cmd.SetArgs([]string{"sync"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("sync command returned error: %v", err)
	}

	for _, student := range students {
		fixture := fixtures[student.Name]
		msg := getRemoteHeadCommitSubject(t, fixture.remoteDir, fixture.branchName)
		if msg != app.InstructorSyncCommitMessage {
			t.Fatalf(
				"remote head message for %q = %q, want %q",
				student.Name,
				msg,
				app.InstructorSyncCommitMessage,
			)
		}
	}
}

func setupRepoWithRemote(t *testing.T, repoDir string, remoteDir string) string {
	t.Helper()

	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		t.Fatalf("MkdirAll repo dir failed: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(remoteDir), 0o755); err != nil {
		t.Fatalf("MkdirAll remote parent dir failed: %v", err)
	}

	runCmd(t, "", "git", "init", "--bare", remoteDir)
	runCmd(t, repoDir, "git", "init")
	runCmd(t, repoDir, "git", "config", "user.name", "Course Sync Test")
	runCmd(t, repoDir, "git", "config", "user.email", "test@example.com")

	notesPath := filepath.Join(repoDir, "notes.txt")
	if err := os.WriteFile(notesPath, []byte("initial notes\n"), 0o644); err != nil {
		t.Fatalf("WriteFile notes failed: %v", err)
	}

	runCmd(t, repoDir, "git", "add", "notes.txt")
	runCmd(t, repoDir, "git", "commit", "-m", "init")
	runCmd(t, repoDir, "git", "remote", "add", "origin", remoteDir)
	runCmd(t, repoDir, "git", "push", "-u", "origin", "HEAD")

	return strings.TrimSpace(runCmd(t, repoDir, "git", "rev-parse", "--abbrev-ref", "HEAD"))
}

func appendFile(t *testing.T, filePath string, content string) {
	t.Helper()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		t.Fatalf("OpenFile failed: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
}

func getRemoteHeadCommitSubject(t *testing.T, remoteDir string, branchName string) string {
	t.Helper()

	refName := fmt.Sprintf("refs/heads/%s", branchName)
	hash := strings.TrimSpace(runCmd(
		t,
		"",
		"git",
		"--git-dir",
		remoteDir,
		"rev-parse",
		refName,
	))

	return strings.TrimSpace(runCmd(
		t,
		"",
		"git",
		"--git-dir",
		remoteDir,
		"log",
		"-1",
		"--pretty=%s",
		hash,
	))
}

func runCmd(t *testing.T, dir string, name string, args ...string) string {
	t.Helper()

	cmd := exec.Command(name, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command failed: %s %v\nerror: %v\noutput:\n%s", name, args, err, string(output))
	}

	return string(output)
}
