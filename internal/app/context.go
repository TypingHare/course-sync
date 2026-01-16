package app

import (
	"os"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/io"
)

// Context represents the application context, including the working directory,
// project directory, and user's role.
type Context struct {
	*io.OutputMode // Output mode for controlling verbosity and styling.

	WorkingDir string // Current working directory.
	ProjectDir string // Project directory.

	Role model.Role // User's role (instructor or student).
}

// NewContext creates a new application context by determining the working
// directory, project directory, and user's role.
func NewContext() (*Context, error) {
	// Determine working directory.
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Determine project directory.
	projectDir, err := FindProjectDir(workingDir)
	if err != nil {
		return nil, err
	}

	// Determine user's role.
	role, err := GetRole(GetDataDir(projectDir))
	if err != nil {
		return nil, err
	}

	return &Context{
		OutputMode: io.NewOutputMode(false, false, false),
		WorkingDir: workingDir,
		ProjectDir: projectDir,
		Role:       role,
	}, nil
}

// IsInstructor returns true if the user has the instructor role, and false
// otherwise.
func (c *Context) IsInstructor() bool {
	return c.Role == model.RoleInstructor
}

// IsStudent returns true if the user has the student role, and false otherwise.
func (c *Context) IsStudent() bool {
	return c.Role == model.RoleStudent
}
