package ssh

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/domain/port"
	"github.com/TypingHare/course-sync/internal/infra/exec"
	"github.com/TypingHare/course-sync/internal/infra/fs"
)

// instructor private key file name inside the application directory.
const InstructorPrivateKeyFileName = "instructor"

// instructor public key file name inside the application directory.
const InstructorPublicKeyFileName = "instructor.pub"

// GenerateKeyPair generates a new SSH key pair. This function ensures that the
// parent directory for the private key file exists before generating the
// pair. Then, it uses the ssh-keygen command to create an ed25519 key pair
// without a passphrase.
func GenerateKeyPair(
	outputMode port.OutputMode,
	projectDir string,
	privateKeyFile string,
) error {
	err := exec.ShellEnsureDir(
		outputMode,
		projectDir,
		filepath.Dir(privateKeyFile),
	)
	if err != nil {
		return fmt.Errorf("ensure parent directory exists: %w", err)
	}

	return exec.NewCommandTask(
		outputMode,
		[]string{
			"ssh-keygen",
			"-t",
			"ed25519",
			"-f",
			privateKeyFile,
			"-N",
			"",
		},
		"Generating new SSH key pair...",
		"Generated SSH key pair successfully.",
		"Failed to generate SSH key pair.",
	).StartE()
}

// GenerateInstructorKeyPair generates a new instructor SSH key pair and saves
// it to the application directory.
func GenerateInstructorKeyPair(
	outputMode port.OutputMode,
	projectDir string,
	appDataDir string,
	force bool,
) error {
	instructorPrivateKeyFile := filepath.Join(
		appDataDir,
		InstructorPrivateKeyFileName,
	)

	instructorPrivateKeyFileExists, err := fs.FileExists(
		instructorPrivateKeyFile,
	)
	if err != nil {
		return err
	}

	if instructorPrivateKeyFileExists && !force {
		return fmt.Errorf(
			"instructor private key file already exists at %s",
			instructorPrivateKeyFile,
		)
	}

	return GenerateKeyPair(outputMode, projectDir, instructorPrivateKeyFile)
}
