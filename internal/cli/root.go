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

func Cmd(ctx *app.Context) *cobra.Command {
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
	cmd.PersistentFlags().BoolVar(
		&ctx.Verbose,
		"verbose", false,
		"enable verbose output",
	)
	cmd.PersistentFlags().BoolVar(
		&ctx.Quiet,
		"quiet", false,
		"suppress non-error output",
	)
	cmd.PersistentFlags().BoolVar(
		&ctx.Plain,
		"plain", false,
		"disable styled output",
	)

	cmd.AddCommand(contextCmd(ctx))
	cmd.AddCommand(userCmd(ctx))
	cmd.AddCommand(pathCmd(ctx))
	cmd.AddCommand(syncCmd(ctx))
	cmd.AddCommand(assignment.Cmd(ctx))
	cmd.AddCommand(doc.Cmd(ctx))
	cmd.AddCommand(grade.Cmd(ctx))
	cmd.AddCommand(submission.Cmd(ctx))
	cmd.AddCommand(ssh.Cmd(ctx))

	if ctx.IsInstructor() {
		cmd.AddCommand(student.Cmd(ctx))
	}

	// Stop Cobra from printing usage or errors automatically.
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	return cmd
}
