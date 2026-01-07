package feature

import "github.com/TypingHare/course-sync/internal/external"

// Pull pulls the latest changes from the remote Git repository.
func Pull() error {
	command := external.ExternalCommand{
		Command:        "git",
		Args:           []string{"pull"},
		OngoingMessage: "Pulling changes from remote repository...",
		DoneMessage:    "Pulled changes from remote repository.",
		ErrorMessage:   "Failed to pull changes from remote repository.",
		PrintCommand:   true,
		PrintStdout:    true,
		PrintStderr:    true,
	}

	return command.Start()
}
