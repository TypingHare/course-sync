package feature

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/execx"
)

// GenerateKeyPair generates a new SSH key pair and saves it to the specified directory.
func GenerateKeyPair(dirPath string) error {
	filePath := filepath.Join(dirPath, app.MASTER_PRIVATE_KEY_FILE_NAME)

	commandTask := execx.CommandTask{
		Command:        "ssh-keygen",
		Args:           []string{"-t", "ed25519", "-f", filePath, "-N", ""},
		OngoingMessage: "Generating new SSH key pair...",
		DoneMessage:    "SSH key pair generated successfully.",
		ErrorMessage:   "Failed to generate SSH key pair.",
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}

// GenerateMasterKeyPair generates a new SSH master key pair and saves it to the application
// directory.
func GenerateMasterKeyPair() error {
	appDirPath, err := app.GetAppDirPath()
	if err != nil {
		return err
	}

	return GenerateKeyPair(appDirPath)
}
