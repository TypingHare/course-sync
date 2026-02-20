package app

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/filesystem"
)

// GetRole determines the current role based on instructor key presence.
func GetRole(dataDir string) (model.Role, error) {
	isInstructor, err := filesystem.FileExists(
		GetInstructorPrivateKeyFile(dataDir),
	)
	if err != nil {
		return model.RoleUnknown, err
	}

	if isInstructor {
		return model.RoleInstructor, nil
	}

	return model.RoleStudent, nil
}
