package shell

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

// RunCommand executes the given command with arguments and returns its standard output. If the
// command exits with a non-zero status or fails to start, an error is returned containing the
// underlying error and any output written to standard error.
func RunCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command failed: %w\n%s", err, stderr.String())
	}

	return stdout.String(), nil
}

func MakeOngoing(message string) string {
	return color.New(color.FgHiCyan).SprintFunc()(message)
}

func MakeDone(message string) string {
	return color.New(color.FgHiBlack).SprintFunc()(message)
}

func MakeSuccess(message string) string {
	return fmt.Sprintf("🟩 %s", message)
}

func MakeWarning(message string) string {
	return fmt.Sprintf("🟧 %s", message)
}

func MakeError(message string) string {
	return fmt.Sprintf("🟥 %s", message)
}
