package app

import (
	"os"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/io"
)

type Context struct {
	*io.OutputMode

	WorkingDir string
	ProjectDir string

	Role model.Role
}

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

func (c *Context) IsInstructor() bool {
	return c.Role == model.RoleInstructor
}

func (c *Context) IsStudent() bool {
	return c.Role == model.RoleStudent
}
