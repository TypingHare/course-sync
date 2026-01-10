package git

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/exec"
)

// GitGetUsername retrieves the Git user name from the Git configuration.
func GitGetUsername(appCtx *app.Context) (string, error) {
	commandTask, err := exec.NewCommandTask(
		appCtx,
		[]string{"git", "config", "--get", "user.name"},
		"Retrieving Git user name...",
		"Retrieved Git user name.",
		"Failed to retrieve Git user name.",
	)
	if err != nil {
		return "", err
	}

	result, err := commandTask.Start()
	if err != nil {
		return "", fmt.Errorf("get git user name: %w", err)
	}

	return strings.TrimSpace(result.Stdout), nil
}
