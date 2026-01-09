package feature

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
)

// Role represents the user role in the application.
type Role string

const (
	StudentRole Role = "student"
	MasterRole  Role = "master"
)

// RoleCache caches the determined role to avoid redundant file system checks.
var RoleCache Role = ""

// GetRole determines the current user role based on the presence of the master private key file.
func GetRole() (Role, error) {
	if RoleCache != "" {
		return RoleCache, nil
	}

	appDirPath, err := app.GetAppDirPath()
	if err != nil {
		return "", err
	}

	isMaster, err := FileExists(filepath.Join(appDirPath, app.MASTER_PRIVATE_KEY_FILE_NAME))
	if err != nil {
		return "", err
	}

	if isMaster {
		RoleCache = MasterRole
	} else {
		RoleCache = StudentRole
	}

	return RoleCache, nil
}
