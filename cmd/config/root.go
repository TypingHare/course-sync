package config

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration settings for Course Sync.",
		Long: `The config command allows you to view and modify configuration settings
for Course Sync, including repository URLs, authentication methods, and
other preferences.`,
	}

	return cmd
}
