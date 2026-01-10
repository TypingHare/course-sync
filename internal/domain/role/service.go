package role

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/domain/ssh"
	"github.com/TypingHare/course-sync/internal/infra/fs"
)

// Role represents the user role in the application.
type Role string

// Defined user roles.
const (
	RoleStudent Role = "student"
	RoleMaster  Role = "master"
)

// GetRole determines the current user role based on the presence of the master
// private key file.
func GetRole(appDir string) (Role, error) {
	isMaster, err := fs.FileExists(
		filepath.Join(appDir, ssh.MASTER_PRIVATE_KEY_FILE_NAME),
	)
	if err != nil {
		return "", err
	}

	if isMaster {
		return RoleMaster, nil
	} else {
		return RoleStudent, nil
	}
}
