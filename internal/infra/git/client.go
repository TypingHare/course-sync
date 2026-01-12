package git

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/domain/port"
	"github.com/TypingHare/course-sync/internal/infra/exec"
)

// Clone clones a Git repository from the specified URL into the destination
// path.
func Clone(outputMode port.OutputMode, repoURL string, destPath string) error {
	return exec.NewCommandTask(
		outputMode,
		[]string{"git", "clone", repoURL, destPath},
		fmt.Sprintf("Cloning repository %q into %q...", repoURL, destPath),
		fmt.Sprintf("Cloned repository %q into %q.", repoURL, destPath),
		fmt.Sprintf(
			"Failed to clone repository %q into %q.",
			repoURL,
			destPath,
		),
	).StartE()
}

// GetUsername retrieves the Git user name from the Git configuration.
func GetUsername(outputMode port.OutputMode) (string, error) {
	commandTask := exec.NewCommandTask(
		outputMode,
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

// Add stages the specified file for commit in Git.
func Add(outputMode port.OutputMode, path string) error {
	return exec.NewCommandTask(
		outputMode,
		[]string{"git", "add", path},
		fmt.Sprintf("Staging file/dir %q for commit...", path),
		fmt.Sprintf("Staged file/dir %q for commit.", path),
		fmt.Sprintf("Failed to stage file/dir %q for commit.", path),
	).StartE()
}

// Commit commits the staged changes to Git with the provided commit message.
func Commit(outputMode port.OutputMode, message string) error {
	return exec.NewCommandTask(
		outputMode,
		[]string{"git", "commit", "-m", message},
		fmt.Sprintf("Committing changes (%s)...", message),
		fmt.Sprintf("Committed changes (%s)...", message),
		fmt.Sprintf("Failed to commit changes (%s).", message),
	).StartE()
}

// Push pushes the local changes to the remote Git repository.
func Push(outputMode port.OutputMode) error {
	return exec.NewCommandTask(
		outputMode,
		[]string{"git", "push"},
		"Pushing changes to remote repository...",
		"Pushed changes to remote repository.",
		"Failed to push changes to remote repository.",
	).StartE()
}

// Pull pulls the latest changes from the remote Git repository.
func Pull(outputMode port.OutputMode, rebase bool) error {
	args := []string{"git", "pull"}
	if rebase {
		args = append(args, "--rebase")
	}

	return exec.NewCommandTask(
		outputMode,
		args,
		"Pulling changes from remote repository...",
		"Pulled changes from remote repository.",
		"Failed to pull changes from remote repository.",
	).StartE()
}

// Restore restores the specified file to its last committed state in Git.
func Restore(outputMode port.OutputMode, path string) error {
	return exec.NewCommandTask(
		outputMode,
		[]string{"git", "restore", path},
		fmt.Sprintf("Restoring %q...", path),
		fmt.Sprintf("Restored %q.", path),
		fmt.Sprintf("Failed to restore %q.", path),
	).StartE()
}

// RevParseHead retrieves the current Git commit hash (HEAD).
// It returns the commit hash as a string.
func RevParseHead(outputMode port.OutputMode) (string, error) {
	commandTask := exec.NewCommandTask(
		outputMode,
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
