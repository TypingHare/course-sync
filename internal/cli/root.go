package cli

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "csync",
		Short:   "Course Sync (csync) helps students and teachers synchronize course materials.",
		Long:    ``,
		Version: "2026.1",
	}

	cmd.SetVersionTemplate("{{.Version}}\n")

	return cmd
}
