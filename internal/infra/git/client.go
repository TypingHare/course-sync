package git

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/exec"
)

// GitGetUsername retrieves the Git user name from the Git configuration.
func GitGetUsername(appCtx *app.Context) (string, error) {
	commandTask := exec.NewCommandTask(
		appCtx,
		[]string{"git", "config", "--get", "user.name"},
		"Retrieving Git user name...",
		"Retrieved Git user name.",
		"Failed to retrieve Git user name.",
	)

	result, err := commandTask.Start()
	if err != nil {
		return "", fmt.Errorf("get git user name: %w", err)
	}

	return strings.TrimSpace(strings.TrimSpace(result.Stdout)), nil
}

// GitAdd stages the specified file for commit in Git.
func GitAdd(appCtx *app.Context, path string) error {
	return exec.NewCommandTask(
		appCtx,
		[]string{"git", "add", path},
		fmt.Sprintf("Staging file/dir %q for commit...", path),
		fmt.Sprintf("Staged file/dir %q for commit.", path),
		fmt.Sprintf("Failed to stage file/dir %q for commit.", path),
	).StartE()
}

// GitCommit commits the staged changes to Git with the provided commit message.
func GitCommit(appCtx *app.Context, message string) error {
	return exec.NewCommandTask(
		appCtx,
		[]string{"git", "commit", "-m", message},
		fmt.Sprintf("Committing changes (%s)...", message),
		fmt.Sprintf("Committed changes (%s)...", message),
		fmt.Sprintf("Failed to commit changes (%s).", message),
	).StartE()
}

// GitPush pushes the local changes to the remote Git repository.
func GitPush(appCtx *app.Context) error {
	return exec.NewCommandTask(
		appCtx,
		[]string{"git", "push"},
		"Pushing changes to remote repository...",
		"Pushed changes to remote repository.",
		"Failed to push changes to remote repository.",
	).StartE()
}

// GitPull pulls the latest changes from the remote Git repository.
func GitPull(appCtx *app.Context) error {
	return exec.NewCommandTask(
		appCtx,
		[]string{"git", "pull"},
		"Pulling changes from remote repository...",
		"Pulled changes from remote repository.",
		"Failed to pull changes from remote repository.",
	).StartE()
}

// GitRevParseHead retrieves the current Git commit hash (HEAD).
// It returns the commit hash as a string.
func GitRevParseHead(appCtx *app.Context) (string, error) {
	commandTask := exec.NewCommandTask(
		appCtx,
		[]string{"git", "rev-parse", "HEAD"},
		"Retrieving current Git commit hash...",
		"Retrieved current Git commit hash.",
		"Failed to retrieve current Git commit hash.",
	)

	result, err := commandTask.Start()
	if err != nil {
		return "", fmt.Errorf("get git commit hash: %w", err)
	}

	return strings.TrimSpace(result.Stdout), nil
}
