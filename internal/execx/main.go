package execx

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

// RunCommand executes an external command and returns its exit code, standard output, standard
// error, and any execution error.
//
// The command is executed directly (without invoking a shell), using the provided name and
// arguments. Standard output and standard error are captured separately and returned regardless of
// whether the command succeeds.
//
// The returned exitCode follows these rules:
//
//   - 0 if the command exits successfully
//   - the process exit status if the command exits with a non-zero status
//   - -1 if the command fails to start or the exit status cannot be determined
//
// If the command exits with a non-zero status or fails to start, the returned error will be non-nil
// and wrap the underlying execution error. Callers may inspect the error (e.g., using errors.As) to
// detect exit status details.
//
// RunCommand does not print output or modify terminal state; callers are responsible for displaying
// or handling the returned output as appropriate.
func RunCommand(
	name string,
	args ...string,
) (exitCode int, stdout string, stderr string, err error) {
	cmd := exec.Command(name, args...)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err = cmd.Run()

	stdout = outBuf.String()
	stderr = errBuf.String()

	if err == nil {
		return 0, stdout, stderr, nil
	}

	// Command started but exited with a non-zero status
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return exitErr.ExitCode(), stdout, stderr,
			fmt.Errorf("command %q failed with exit code %d: %w", name, exitErr.ExitCode(), err)
	}

	// Command failed to start or exit code is unavailable
	return -1, stdout, stderr,
		fmt.Errorf("command %q failed to start: %w", name, err)
}
