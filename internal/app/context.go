package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/domain/role"
	"github.com/TypingHare/course-sync/internal/infra/config"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"

	"github.com/TypingHare/course-sync/internal/infra/fs"
)

// Configuration file name.
const CONFIG_FILE_NAME = "config.json"

// Application data directory name.
const APP_DATA_DIR_NAME = ".csync"

// SRC_DIR_NAME is the name of the source directory within the project.
const SRC_DIR_NAME = "src"

// DOCS_FILE_NAME is the name of the documentation file.
const DOCS_DIR_NAME = "docs"

// Master private key file name inside the application directory.
const MASTER_PRIVATE_KEY_FILE_NAME = "master"

// Master public key file name inside the application directory.
const MASTER_PUBLIC_KEY_FILE_NAME = "master.pub"

// Context holds the context for the Course Sync application.
type Context struct {
	// CLI options (flags).
	Verbose bool // Enable verbose output.
	Quiet   bool // Enable quiet mode, suppressing output.

	// Environment settings.
	WorkingDir string    // Current working directory.
	ProjectDir string    // Absolute path to the project root directory.
	AppDataDir string    // Absolute path to the application data directory.
	SrcDir     string    // Absolute path to the source directory.
	DocsDir    string    // Absolute path to the documentation directory.
	Role       role.Role // User role in the application.

	// Application configuration.
	config *config.Config // Application configuration.
}

// BuildContext initializes and returns the application context.
func BuildContext() (*Context, error) {
	ctx := &Context{}

	// Resolve working directory.
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	ctx.WorkingDir = workingDir

	// Find project directory.
	projectDir, err := FindProjectDir(workingDir)
	if err != nil {
		return nil, err
	}
	ctx.ProjectDir = projectDir

	// Set application data directory path.
	ctx.AppDataDir = filepath.Join(
		ctx.ProjectDir,
		APP_DATA_DIR_NAME,
	)

	// Set source directory path.
	ctx.SrcDir = filepath.Join(
		ctx.ProjectDir,
		SRC_DIR_NAME,
	)

	// Set docs directory path.
	ctx.DocsDir = filepath.Join(
		ctx.ProjectDir,
		DOCS_DIR_NAME,
	)

	// Determine user role.
	userRole, err := GetRole(ctx.ProjectDir)
	if err != nil {
		return nil, err
	}
	ctx.Role = userRole

	return ctx, nil
}

// GetRole determines the current user role based on the presence of the master
// private key file, which is stored in the application data directory.
func GetRole(appDataDir string) (role.Role, error) {
	isMaster, err := fs.FileExists(
		filepath.Join(appDataDir, MASTER_PRIVATE_KEY_FILE_NAME),
	)
	if err != nil {
		return "", err
	}

	if isMaster {
		return role.RoleMaster, nil
	} else {
		return role.RoleStudent, nil
	}
}

// GetRelPath converts an absolute path to a path relative to the project
// directory.
func (ctx *Context) GetRelPath(absPath string) (string, error) {
	relPath, err := filepath.Rel(ctx.ProjectDir, absPath)
	if err != nil {
		return "", fmt.Errorf("can't make %q to %q", absPath, ctx.ProjectDir)
	}

	return relPath, nil
}

// GetConfig retrieves the application configuration, loading it from the Config
// file if necessary.
func (ctx *Context) GetConfig() *config.Config {
	if ctx.config == nil {
		config, _ := jsonstore.ReadJSONFile[config.Config](
			filepath.Join(ctx.ProjectDir, CONFIG_FILE_NAME),
		)
		ctx.config = &config
	}

	return ctx.config
}

// SaveConfig saves the current configuration to the config file.
func (ctx *Context) SaveConfig() error {
	return jsonstore.WriteJSONFile(CONFIG_FILE_NAME, ctx.config)
}
