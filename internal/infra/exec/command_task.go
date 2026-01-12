package exec

import (
	"fmt"
	"os"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/ui"
	"github.com/fatih/color"
)

// CommandTask represents a task that runs an external command with specified
// arguments and displays messages based on its progress and outcome.
type CommandTask struct {
	// The external command to run.
	command string

	// Arguments for the external command.
	args []string

	// Message to display while the command is running.
	ongoingMessage string

	// Message to display when the command completes successfully.
	doneMessage string

	// Message to display if the command fails.
	errorMessage string

	// Whether to print the command being executed.
	verbose bool

	// Whether to print the command's standard output.
	quiet bool
}

// CommandTaskResult holds the result of executing a CommandTask.
type CommandTaskResult struct {
	ExitCode int    // Exit code of the command.
	Stdout   string // Captured standard output.
	Stderr   string // Captured standard error.
}

// NewCommandTask creates a new CommandTask with the given parameters.
func NewCommandTask(
	appCtx *app.Context,
	args []string,
	ongoingMessage string,
	doneMessage string,
	errorMessage string,
) *CommandTask {
	var command string
	if len(args) == 0 {
		command = ""
	} else {
		command = args[0]
	}

	return &CommandTask{
		command:        command,
		args:           args[1:],
		ongoingMessage: ongoingMessage,
		doneMessage:    doneMessage,
		errorMessage:   errorMessage,
		verbose:        appCtx.Verbose,
		quiet:          appCtx.Quiet,
	}
}

// Start executes the command task, displaying messages based on its progress
// and outcome. It returns the result of the command execution and any error
// encountered.
func (t *CommandTask) Start() (CommandTaskResult, error) {
	var appendedCommandStr string

	if t.verbose {
		commandStr := JoinCommand(t.command, t.args)
		appendedCommandStr = fmt.Sprintf("\n   [ %s ]", commandStr)
	}

	// Create and start the spinner.
	ongoingMessage := t.ongoingMessage + appendedCommandStr
	spinner := ui.NewSpinner(os.Stdout, ui.MakeOngoing(ongoingMessage))
	if !t.quiet {
		spinner.Start()
	}

	// Execute the external command.
	exitCode, stdout, stderr, err := RunCommand(
		t.command,
		t.args...,
	)

	// Stop the spinner and clear ongoing message before printing results.
	if !t.quiet {
		spinner.Stop()
		spinner.ClearMessage()
	}

	// Print the final message based on success or failure.
	if !t.quiet {
		if err != nil {
			errorMessage := t.errorMessage + appendedCommandStr
			fmt.Println(ui.MarkError(ui.MakeDone(errorMessage)))
		} else {
			doneMessage := t.doneMessage + appendedCommandStr
			fmt.Println(ui.MarkSuccess(ui.MakeDone(doneMessage)))
		}
	}

	// Print stdout and stderr if in verbose mode.
	if !t.quiet && t.verbose {
		PrintExternalCommandStdout(stdout, strings.Repeat(" ", 3))
		PrintExternalCommandStderr(stderr, strings.Repeat(" ", 3))
	}

	return CommandTaskResult{
		ExitCode: exitCode,
		Stdout:   stdout,
		Stderr:   stderr,
	}, err
}

// PrintExternalCommandStdout splits the given stdout string by new lines and
// prints each line indented.
func PrintExternalCommandStdout(stdout string, indentation string) {
	if stdout == "" {
		return
	}

	stdout = strings.TrimRight(stdout, "\n")

	fmt.Println()
	stdoutColor := color.New(color.FgHiYellow).SprintFunc()
	for line := range strings.SplitSeq(stdout, "\n") {
		fmt.Println(indentation + stdoutColor(line))
	}
	fmt.Println()
}

// PrintExternalCommandStderr splits the given stderr string by new lines and
// prints each line indented.
func PrintExternalCommandStderr(stderr string, indentation string) {
	if stderr == "" {
		return
	}

	stderr = strings.TrimRight(stderr, "\n")

	fmt.Println()
	stderrColor := color.New(color.FgHiRed).SprintFunc()
	for line := range strings.SplitSeq(stderr, "\n") {
		fmt.Println(indentation + stderrColor(line))
	}
	fmt.Println()
}

// StartE is a convenience method that starts the command task and returns only
// the error, ignoring the result.
func (t *CommandTask) StartE() error {
	_, err := t.Start()

	return err
}
