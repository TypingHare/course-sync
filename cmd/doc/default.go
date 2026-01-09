package doc

import "github.com/spf13/cobra"

var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Display the default documentation name.",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("default.md")
	},
}
