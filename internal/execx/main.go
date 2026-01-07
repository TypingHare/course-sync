package execx

import (
	"bytes"
	"fmt"
	"os/exec"
)

// RunCommand executes an external command and returns its standard output, standard error, and any
// execution error.
//
// The command is executed directly (without invoking a shell), using the provided name and
// arguments. Standard output and standard error are captured separately and returned regardless of
// whether the command succeeds.
//
// If the command exits with a non-zero status or fails to start, the returned error will be non-nil
// and wrap the underlying execution error. Callers may inspect the error (e.g., using errors.As) to
// detect exit status details.
//
// RunCommand does not print output or modify terminal state; callers are responsible for displaying
// or handling the returned output as appropriate.
func RunCommand(name string, args ...string) (string, string, error) {
	cmd := exec.Command(name, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(),
			fmt.Errorf("command %q failed: %w", name, err)
	}

	return stdout.String(), stderr.String(), nil
}
