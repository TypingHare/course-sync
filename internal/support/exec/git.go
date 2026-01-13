package exec

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/support/io"
)

// GetUsername retrieves the Git user name from the Git configuration.
func GitGetUsername(outputMode *io.OutputMode) (string, error) {
	commandRunner := NewCommandRunner(
		outputMode,
		[]string{"git", "config", "--get", "user.name"},
		"Retrieving Git user name...",
		"Retrieved Git user name.",
		"Failed to retrieve Git user name.",
	)

	result, err := commandRunner.Start()
	if err != nil {
		return "", fmt.Errorf("(git) get user name: %w", err)
	}

	return strings.TrimSpace(strings.TrimSpace(result.Stdout)), nil
}

// GitAdd stages the specified file for commit in Git.
func GitAdd(outputMode *io.OutputMode, path string) error {
	return NewCommandRunner(
		outputMode,
		[]string{"git", "add", path},
		fmt.Sprintf("Staging file/dir %q for commit...", path),
		fmt.Sprintf("Staged file/dir %q for commit.", path),
		fmt.Sprintf("Failed to stage file/dir %q for commit.", path),
	).StartE()
}

// GitCommit commits the staged changes to Git with the provided commit message.
func GitCommit(outputMode *io.OutputMode, message string) error {
	return NewCommandRunner(
		outputMode,
		[]string{"git", "commit", "-m", message},
		fmt.Sprintf("Committing changes (%s)...", message),
		fmt.Sprintf("Committed changes (%s)...", message),
		fmt.Sprintf("Failed to commit changes (%s).", message),
	).StartE()
}

// GitPush pushes the local changes to the remote Git repository.
func GitPush(outputMode *io.OutputMode) error {
	return NewCommandRunner(
		outputMode,
		[]string{"git", "push"},
		"Pushing changes to remote repository...",
		"Pushed changes to remote repository.",
		"Failed to push changes to remote repository.",
	).StartE()
}

// GitPull pulls the latest changes from the remote Git repository.
func GitPull(outputMode *io.OutputMode, rebase bool) error {
	args := []string{"git", "pull"}
	if rebase {
		args = append(args, "--rebase")
	}

	return NewCommandRunner(
		outputMode,
		args,
		"Pulling changes from remote repository...",
		"Pulled changes from remote repository.",
		"Failed to pull changes from remote repository.",
	).StartE()
}

// GitRestore restores the specified file to its last committed state in Git.
func GitRestore(outputMode *io.OutputMode, path string) error {
	return NewCommandRunner(
		outputMode,
		[]string{"git", "restore", path},
		fmt.Sprintf("Restoring %q...", path),
		fmt.Sprintf("Restored %q.", path),
		fmt.Sprintf("Failed to restore %q.", path),
	).StartE()
}

// GitRevParseHead retrieves the current Git commit hash (HEAD).
// It returns the commit hash as a string.
func GitRevParseHead(outputMode *io.OutputMode) (string, error) {
	commandTask := NewCommandRunner(
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
