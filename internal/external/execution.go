package external

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/ui"
	"github.com/fatih/color"
)

type ExternalCommand struct {
	Command        string   // The external command to run.
	Args           []string // Arguments for the external command.
	OngoingMessage string   // Message to display while the command is running.
	DoneMessage    string   // Message to display when the command completes successfully.
	ErrorMessage   string   // Message to display if the command fails.
	PrintCommand   bool     // Whether to print the external command.
	PrintStdout    bool     // Whether to print the standard output of the command.
	PrintStderr    bool     // Whether to print the standard error of the command.
}

func (externalCommand *ExternalCommand) Start() error {
	ongoingMessage := externalCommand.OngoingMessage
	errorMessage := externalCommand.ErrorMessage
	doneMessage := externalCommand.DoneMessage

	if externalCommand.PrintCommand {
		commandStr := externalCommand.Command + " " + strings.Join(externalCommand.Args, " ")
		ongoingMessage += fmt.Sprintf(" (%s)", commandStr)
		doneMessage += fmt.Sprintf(" (%s)", commandStr)
		errorMessage += fmt.Sprintf(" (%s)", commandStr)
	}
	spinner := ui.NewSpinner(ui.MakeOngoing(ongoingMessage))
	spinner.Start()

	stdout, stderr, err := RunCommand(externalCommand.Command, externalCommand.Args...)
	spinner.Stop()
	if err != nil {
		fmt.Println(ui.MarkError(errorMessage))
	} else {
		fmt.Println(ui.MarkSuccess(errorMessage))
	}

	if externalCommand.PrintStdout {
		PrintExternalCommandStdout(stdout, strings.Repeat(" ", 3))
	}

	if externalCommand.PrintStderr {
		PrintExternalCommandStderr(stderr, strings.Repeat(" ", 3))
	}

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
	stdoutColor := color.New(color.FgHiYellow).SprintFunc()
	for line := range strings.SplitSeq(stderr, "\n") {
		fmt.Println(indentation + stdoutColor(line))
	}
	fmt.Println()
}
