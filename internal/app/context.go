package app

import (
	"os"

	"github.com/TypingHare/course-sync/internal/infra/config"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

// Configuration file name.
const CONFIG_FILE_NAME = "config.json"

// Singleton application context.
var appContext *AppContext

// AppContext holds the context for the Course Sync application.
type AppContext struct {
	// CLI options (flags).
	Verbose bool // Enable verbose output.
	Quiet   bool // Enable quiet mode, suppressing output.

	// Environment settings.
	WorkingDir string         // Current working directory.
	ProjectDir string         // Path to the project root directory.
	Config     *config.Config // Application configuration.
}

// BuildContext initializes and returns the application context.
func BuildContext() (*AppContext, error) {
	appContext = &AppContext{}

	// Resolve working directory.
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	appContext.WorkingDir = workingDir

	// Find project directory.
	projectDir, err := FindProjectDir(workingDir)
	if err != nil {
		return nil, err
	}
	appContext.ProjectDir = projectDir

	// Load configuration.
	config, _ := jsonstore.ReadJSONFile[config.Config](CONFIG_FILE_NAME)
	appContext.Config = &config

	return appContext, nil
}

// SaveConfig saves the current configuration to the config file.
func (ctx *AppContext) SaveConfig() error {
	return jsonstore.WriteJSONFile(CONFIG_FILE_NAME, ctx.Config)
}
