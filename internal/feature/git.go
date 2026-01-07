package feature

import "github.com/TypingHare/course-sync/internal/execx"

// Pull pulls the latest changes from the remote Git repository.
func Pull(verbose bool) error {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"pull"},
		OngoingMessage: "Pulling changes from remote repository...",
		DoneMessage:    "Pulled changes from remote repository.",
		ErrorMessage:   "Failed to pull changes from remote repository.",
		PrintCommand:   verbose,
		PrintStdout:    verbose,
		PrintStderr:    verbose,
	}

	return commandTask.Start()
}

// Push pushes the local changes to the remote Git repository.
func Push(verbose bool) error {
	commandTask := execx.CommandTask{
		Command:        "git",
		Args:           []string{"push"},
		OngoingMessage: "Pushing changes to remote repository...",
		DoneMessage:    "Pushed changes to remote repository.",
		ErrorMessage:   "Failed to push changes to remote repository.",
		PrintCommand:   verbose,
		PrintStdout:    verbose,
		PrintStderr:    verbose,
	}

	return commandTask.Start()
}
