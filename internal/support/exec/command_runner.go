package exec

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TypingHare/course-sync/internal/support/io"
	"github.com/fatih/color"
)

// CommandRunner represents a task that runs an external command with specified
// arguments and displays messages based on its progress and outcome.
type CommandRunner struct {
	// The working directory for the command execution.
	workingDir string

	// outputMode is the output mode for controlling verbosity and quietness.
	outputMode *io.OutputMode

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
}

// CommandRunnerResult holds the result of executing a CommandRunner.
type CommandRunnerResult struct {
	ExitCode int    // Exit code of the command.
	Stdout   string // Captured standard output.
	Stderr   string // Captured standard error.
}

// NewCommandRunner creates a new CommandRunner with the given parameters.
func NewCommandRunner(
	outputMode *io.OutputMode,
	args []string,
	ongoingMessage string,
	doneMessage string,
	errorMessage string,
) *CommandRunner {
	var command string
	if len(args) == 0 {
		command = ""
	} else {
		command = args[0]
	}

	return &CommandRunner{
		workingDir:     "",
		outputMode:     outputMode,
		command:        command,
		args:           args[1:],
		ongoingMessage: ongoingMessage,
		doneMessage:    doneMessage,
		errorMessage:   errorMessage,
	}
}

// SetWorkingDir sets the working directory for the command runner and returns
// the updated CommandRunner.
func (t *CommandRunner) SetWorkingDir(dir string) *CommandRunner {
	t.workingDir = dir
	return t
}

// Start executes the command runner, displaying messages based on its progress
// and outcome. It returns the result of the command execution and any error
// encountered.
func (t *CommandRunner) Start() (CommandRunnerResult, error) {
	var appendedCommandStr string

	if t.outputMode.IsVerbose() {
		commandStr := JoinCommand(t.command, t.args)
		appendedCommandStr = fmt.Sprintf("\n   [ %s ]", commandStr)
	}

	// Create and start the spinner.
	ongoingMessage := t.ongoingMessage + appendedCommandStr
	spinner := io.NewSpinner(os.Stdout, io.MakeOngoing(ongoingMessage))
	if !t.outputMode.IsQuiet() {
		spinner.Start()
	}

	// This is made intentional to allow students to read the ongoing message
	// for a bit longer in verbose mode.
	if t.outputMode.IsVerbose() {
		time.Sleep(io.SpinnerFrameTime * 4)
	}

	// Execute the external command.
	exitCode, stdout, stderr, err := RunCommand(
		t.workingDir,
		t.command,
		t.args...,
	)

	// Stop the spinner and clear ongoing message before printing results.
	if !t.outputMode.IsQuiet() {
		spinner.Stop()
		spinner.ClearMessage()
	}

	// Print the final message based on success or failure.
	if !t.outputMode.IsQuiet() {
		if err != nil {
			errorMessage := t.errorMessage + appendedCommandStr
			fmt.Println(io.MarkError(io.MakeDone(errorMessage)))
		} else {
			doneMessage := t.doneMessage + appendedCommandStr
			fmt.Println(io.MarkSuccess(io.MakeDone(doneMessage)))
		}
	}

	// Print stdout and stderr if in verbose mode.
	if !t.outputMode.IsQuiet() && t.outputMode.IsVerbose() {
		PrintExternalCommandStdout(stdout, strings.Repeat(" ", 3))
		PrintExternalCommandStderr(stderr, strings.Repeat(" ", 3))
	}

	return CommandRunnerResult{
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

// StartE is a convenience method that starts the command runner and returns
// only the error, ignoring the result.
func (t *CommandRunner) StartE() error {
	_, err := t.Start()

	return err
}
