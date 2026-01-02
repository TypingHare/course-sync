package feature

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/shell"
	"github.com/fatih/color"
)

type SpinnerExecution struct {
	ongoingMessage string
	doneMessage    string
	errorMessage   string
	command        string
	args           []string
	printCommand   bool
	printStdout    bool
}

func (spinnerExecution SpinnerExecution) Start() error {
	ongoingMessage := spinnerExecution.ongoingMessage
	errorMessage := spinnerExecution.errorMessage
	doneMessage := spinnerExecution.doneMessage
	if spinnerExecution.printCommand {
		commandStr := spinnerExecution.command + " " + strings.Join(spinnerExecution.args, " ")
		ongoingMessage += fmt.Sprintf(" (%s)", commandStr)
		doneMessage += fmt.Sprintf(" (%s)", commandStr)
		errorMessage += fmt.Sprintf(" (%s)", commandStr)
	}
	spinner := shell.NewSpinner(shell.MakeOngoing(ongoingMessage))
	spinner.Start()

	stdout, err := shell.RunCommand(spinnerExecution.command, spinnerExecution.args...)
	if err != nil {
		spinner.StopWithMessage(shell.MakeError(errorMessage))
		fmt.Println("Error:", err.Error())
		PrintShellCommandStdout(stdout, "   ")

		return err
	}

	spinner.StopWithMessage(shell.MakeSuccess(shell.MakeDone(doneMessage)))

	PrintShellCommandStdout(stdout, "   ")

	return nil
}

// PrintShellCommandStdout splits the given stdout string by new lines and prints
// each line indented.
func PrintShellCommandStdout(stdout string, indentation string) {
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
