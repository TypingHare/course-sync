package feature

import (
	"errors"

	"github.com/TypingHare/course-sync/internal/app"
)

// MakeSyncCommit stages all changes in the application directory and creates a commit with the
// message "sync".
func MakeSyncCommit() error {
	appDir, err := app.GetAppDirPath()
	if err != nil {
		return errors.New("failed to get application directory path: " + err.Error())
	}

	return app.RunAll(
		func() error { return GitAdd(appDir) },
		func() error { return GitCommit("sync") },
	)
}
