package feature

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/execx"
)

// GetGitUserName retrieves the Git user name from the Git configuration.
func GetGitUserName() (string, error) {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"config", "--get", "user.name"},
		OngoingMessage: "Retrieving Git user name...",
		DoneMessage:    "Retrieved Git user name.",
		ErrorMessage:   "Failed to retrieve Git user name.",
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	err := commandTask.Start()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(commandTask.Stdout), nil
}

// GitAdd stages the specified file for commit in Git.
func GitAdd(filePath string) error {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"add", filePath},
		OngoingMessage: "Adding file to Git staging area...",
		DoneMessage:    "File added to Git staging area.",
		ErrorMessage:   "Failed to add file to Git staging area.",
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}

// Pull pulls the latest changes from the remote Git repository.
func Pull() error {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"pull"},
		OngoingMessage: "Pulling changes from remote repository...",
		DoneMessage:    "Pulled changes from remote repository.",
		ErrorMessage:   "Failed to pull changes from remote repository.",
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}

// Push pushes the local changes to the remote Git repository.
func Push() error {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"push"},
		OngoingMessage: "Pushing changes to remote repository...",
		DoneMessage:    "Pushed changes to remote repository.",
		ErrorMessage:   "Failed to push changes to remote repository.",
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}
