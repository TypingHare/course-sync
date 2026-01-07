package execx

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/ui"
	"github.com/fatih/color"
)

type CommandTask struct {
	Command        string   // The external command to run.
	Args           []string // Arguments for the external command.
	OngoingMessage string   // Message to display while the command is running.
	DoneMessage    string   // Message to display when the command completes successfully.
	ErrorMessage   string   // Message to display if the command fails.
	PrintCommand   bool     // Whether to print the external command.
	PrintStdout    bool     // Whether to print the standard output of the command.
	PrintStderr    bool     // Whether to print the standard error of the command.

	ExitCode int    // Exit code of the command.
	Stdout   string // Captured standard output.
	Stderr   string // Captured standard error.
}

func (commandTask *CommandTask) Start() error {
	ongoingMessage := commandTask.OngoingMessage
	errorMessage := commandTask.ErrorMessage
	doneMessage := commandTask.DoneMessage

	if commandTask.PrintCommand {
		commandStr := commandTask.Command + " " + strings.Join(commandTask.Args, " ")
		appendedCommandStr := fmt.Sprintf("\n   [ %s ]", commandStr)
		ongoingMessage += appendedCommandStr
		doneMessage += appendedCommandStr
		errorMessage += appendedCommandStr
	}

	spinner := ui.NewSpinner(ui.MakeOngoing(ongoingMessage))
	spinner.Start()

	exitCode, stdout, stderr, err := RunCommand(commandTask.Command, commandTask.Args...)

	spinner.Stop()
	if err != nil {
		fmt.Println(ui.MarkError(ui.MakeDone(errorMessage)))
	} else {
		fmt.Println(ui.MarkSuccess(ui.MakeDone(doneMessage)))
	}

	if commandTask.PrintStdout {
		PrintExternalCommandStdout(stdout, strings.Repeat(" ", 3))
	}

	if commandTask.PrintStderr {
		PrintExternalCommandStderr(stderr, strings.Repeat(" ", 3))
	}

	commandTask.ExitCode = exitCode
	commandTask.Stdout = stdout
	commandTask.Stderr = stderr

	return err
}

// PrintExternalCommandStdout splits the given stdout string by new lines and prints
// each line indented.
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
