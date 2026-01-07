package feature

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/execx"
)

func GetGitUserName(quiet bool, verbose bool) (string, error) {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"config", "--get", "user.name"},
		OngoingMessage: "Retrieving Git user name...",
		DoneMessage:    "Retrieved Git user name.",
		ErrorMessage:   "Failed to retrieve Git user name.",
		Quiet:          quiet,
		PrintCommand:   verbose,
		PrintStdout:    verbose,
		PrintStderr:    verbose,
	}

	err := commandTask.Start()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(commandTask.Stdout), nil
}

// Pull pulls the latest changes from the remote Git repository.
func Pull(quiet bool, verbose bool) error {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"pull"},
		OngoingMessage: "Pulling changes from remote repository...",
		DoneMessage:    "Pulled changes from remote repository.",
		ErrorMessage:   "Failed to pull changes from remote repository.",
		Quiet:          quiet,
		PrintCommand:   verbose,
		PrintStdout:    verbose,
		PrintStderr:    verbose,
	}

	return commandTask.Start()
}

// Push pushes the local changes to the remote Git repository.
func Push(quiet bool, verbose bool) error {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"push"},
		OngoingMessage: "Pushing changes to remote repository...",
		DoneMessage:    "Pushed changes to remote repository.",
		ErrorMessage:   "Failed to push changes to remote repository.",
		Quiet:          quiet,
		PrintCommand:   verbose,
		PrintStdout:    verbose,
		PrintStderr:    verbose,
	}

	return commandTask.Start()
}
