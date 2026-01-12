package cli

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli/assignment"
	"github.com/TypingHare/course-sync/internal/cli/doc"
	"github.com/TypingHare/course-sync/internal/cli/grade"
	"github.com/TypingHare/course-sync/internal/cli/ssh"
	"github.com/TypingHare/course-sync/internal/cli/student"
	"github.com/TypingHare/course-sync/internal/cli/submission"
	"github.com/spf13/cobra"
)

func Cmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "csync",
		Short: "Course Sync (csync)",
		Long: strings.TrimSpace(`
Course Sync helps students and teachers synchronize course materials using Git,
without requiring a central server.

The tool is designed to be fast and responsive, leveraging Go’s concurrency
model to efficiently perform filesystem operations and external command
execution.

Course Sync relies heavily on existing POSIX tools and Git commands to implement
its functionality. This design keeps the system simple and transparent, with
minimal hidden behavior. Users can enable the --verbose flag to inspect the
external commands executed under the hood, along with their standard output
and standard error.

Teachers can publish assignments and release course materials directly through
a Git repository, while students can pull updates and submit assignments without
any server-side infrastructure.

Future versions will incorporate asymmetric encryption to protect application-
generated files and prevent tampering with submissions and metadata.
        `),
		Version: "2026.1",
	}

	cmd.SetVersionTemplate("{{.Version}}\n")
	cmd.PersistentFlags().BoolVarP(
		&appCtx.Verbose,
		"verbose", "v", false,
		"enable verbose output",
	)
	cmd.PersistentFlags().BoolVarP(
		&appCtx.Quiet,
		"quiet", "q", false,
		"suppress non-error output",
	)

	cmd.AddCommand(
		contextCmd(appCtx),
		pathCmd(appCtx),
		initCmd(appCtx),
		ssh.Cmd(appCtx),
		doc.Cmd(appCtx),
		assignment.Cmd(appCtx),
		submission.Cmd(appCtx),
		grade.Cmd(appCtx),
	)

	if appCtx.Isinstructor() {
		cmd.AddCommand(student.Cmd(appCtx))
	}

	if appCtx.IsStudent() {
		cmd.AddCommand(
			pullCmd(appCtx),
			commitCmd(appCtx),
			pushCmd(appCtx),
			syncCommand(appCtx),
		)
	}

	// Stop Cobra from handling errors.
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	return cmd
}
