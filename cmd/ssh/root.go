package ssh

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	sshCmd := &cobra.Command{
		Use:   "ssh",
		Short: "Manage SSH keys for secure communication with remote servers",
		Long: `The ssh command allows you to manage SSH keys for secure communication with remote 
servers. You can generate, list, and delete SSH keys used for authentication.`,

		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	sshCmd.AddCommand(keygenCmd)

	return sshCmd
}
