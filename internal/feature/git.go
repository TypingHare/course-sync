package feature

// Pull pulls the latest changes from the remote Git repository.
func Pull() error {
	return SpinnerExecution{
		ongoingMessage: "Pulling changes from remote repository...",
		doneMessage:    "Pulled changes from remote repository.",
		errorMessage:   "Failed to pull changes from remote repository.",
		command:        "git",
		args:           []string{"pull"},
		printCommand:   true,
		printStdout:    true,
	}.Start()
}
