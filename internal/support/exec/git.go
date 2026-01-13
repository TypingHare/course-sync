package exec

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/support/io"
)

// GetUsername retrieves the Git user name from the Git configuration.
func GitGetUsername(outputMode *io.OutputMode) (string, error) {
	commandRunner := NewCommandRunner(
		outputMode,
		[]string{"git", "config", "--get", "user.name"},
		"Retrieving Git user name...",
		"Retrieved Git user name.",
		"Failed to retrieve Git user name.",
	)

	result, err := commandRunner.Start()
	if err != nil {
		return "", fmt.Errorf("(git) get user name: %w", err)
	}

	return strings.TrimSpace(strings.TrimSpace(result.Stdout)), nil
}
