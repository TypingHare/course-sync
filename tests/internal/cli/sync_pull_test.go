package cli_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func TestSyncPullStudentUsesRebase(t *testing.T) {
	projectDir := t.TempDir()
	remoteDir := filepath.Join(projectDir, "remote.git")
	mainDir := filepath.Join(projectDir, "main")
	peerDir := filepath.Join(projectDir, "peer")

	runCmd(t, "", "git", "init", "--bare", remoteDir)
	runCmd(t, "", "git", "clone", remoteDir, mainDir)
	runCmd(t, mainDir, "git", "config", "user.name", "Student")
	runCmd(t, mainDir, "git", "config", "user.email", "student@example.com")

	baseFile := filepath.Join(mainDir, "base.txt")
	if err := os.WriteFile(baseFile, []byte("base\n"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	runCmd(t, mainDir, "git", "add", "base.txt")
	runCmd(t, mainDir, "git", "commit", "-m", "base")
	runCmd(t, mainDir, "git", "push", "-u", "origin", "HEAD")

	// Local (student) commit that has not been pushed yet.
	localFile := filepath.Join(mainDir, "local.txt")
	if err := os.WriteFile(localFile, []byte("local change\n"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	runCmd(t, mainDir, "git", "add", "local.txt")
	runCmd(t, mainDir, "git", "commit", "-m", "local")

	// Remote commit from a peer to create divergence.
	runCmd(t, "", "git", "clone", remoteDir, peerDir)
	runCmd(t, peerDir, "git", "config", "user.name", "Peer")
	runCmd(t, peerDir, "git", "config", "user.email", "peer@example.com")
	peerFile := filepath.Join(peerDir, "remote.txt")
	if err := os.WriteFile(peerFile, []byte("remote change\n"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	runCmd(t, peerDir, "git", "add", "remote.txt")
	runCmd(t, peerDir, "git", "commit", "-m", "remote")
	runCmd(t, peerDir, "git", "push", "origin", "HEAD")

	oldWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd failed: %v", err)
	}
	if err := os.Chdir(mainDir); err != nil {
		t.Fatalf("Chdir failed: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(oldWD)
	})

	ctx := &app.Context{
		OutputMode: io.NewOutputMode(false, true, true),
		ProjectDir: mainDir,
		WorkingDir: mainDir,
		Role:       model.RoleStudent,
	}
	cmd := cli.Cmd(ctx)
	cmd.SetArgs([]string{"sync", "--pull"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("sync --pull returned error: %v", err)
	}

	log2 := strings.TrimSpace(runCmd(t, mainDir, "git", "log", "--pretty=%s", "-2"))
	lines := strings.Split(log2, "\n")
	if len(lines) < 2 {
		t.Fatalf("unexpected log output: %q", log2)
	}
	if lines[0] != "local" || lines[1] != "remote" {
		t.Fatalf("top commits = %q, want local then remote (rebase result)", log2)
	}
}
