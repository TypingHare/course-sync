package app

import (
	"os"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/domain/role"
	"github.com/TypingHare/course-sync/internal/infra/config"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

// Configuration file name.
const CONFIG_FILE_NAME = "config.json"

// Hidden directory name.
const APP_DATA_DIR_NAME = ".csync"

// SRC_DIR_NAME is the name of the source directory within the project.
const SRC_DIR_NAME = "src"

// Singleton application context.
var context *Context

// Context holds the context for the Course Sync application.
type Context struct {
	// CLI options (flags).
	Verbose bool // Enable verbose output.
	Quiet   bool // Enable quiet mode, suppressing output.

	// Environment settings.
	WorkingDir string    // Current working directory.
	ProjectDir string    // Path to the project root directory.
	AppDataDir string    // Path to the application data directory.
	SrcDir     string    // Path to the source directory.
	Role       role.Role // User role in the application.

	// Application configuration.
	config *config.Config // Application configuration.
}

// BuildContext initializes and returns the application context.
func BuildContext() (*Context, error) {
	context = &Context{}

	// Resolve working directory.
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	context.WorkingDir = workingDir

	// Find project directory.
	projectDir, err := FindProjectDir(workingDir)
	if err != nil {
		return nil, err
	}
	context.ProjectDir = projectDir

	// Set application data directory path.
	context.AppDataDir = filepath.Join(
		context.ProjectDir,
		APP_DATA_DIR_NAME,
	)

	// Set source directory path.
	context.SrcDir = filepath.Join(
		context.ProjectDir,
		SRC_DIR_NAME,
	)

	// Determine user role.
	userRole, err := role.GetRole(context.ProjectDir)
	if err != nil {
		return nil, err
	}
	context.Role = userRole

	return context, nil
}

// GetConfig retrieves the application configuration, loading it from the Config
// file if necessary.
func (ctx *Context) GetConfig() *config.Config {
	if ctx.config == nil {
		config, _ := jsonstore.ReadJSONFile[config.Config](CONFIG_FILE_NAME)
		context.config = &config
	}

	return ctx.config
}

// SaveConfig saves the current configuration to the config file.
func (ctx *Context) SaveConfig() error {
	return jsonstore.WriteJSONFile(CONFIG_FILE_NAME, ctx.config)
}
