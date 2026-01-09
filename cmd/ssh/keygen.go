package ssh

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generate a new SSH key pair",
	Long: `The keygen command allows you to generate a new SSH key pair for secure communication 
with remote servers. You can specify the type of key and the file location to save the keys.`,

	Run: func(cmd *cobra.Command, args []string) {
		feature.GenerateMasterKeyPair()
	},
}
