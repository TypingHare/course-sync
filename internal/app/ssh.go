package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/filesystem"
	"github.com/TypingHare/course-sync/internal/support/io"
)

// GenerateKeyPair generates a new SSH key pair. This function ensures that the
// parent directory for the private key file exists before generating the
// pair. Then, it uses the ssh-keygen command to create an ed25519 key pair
// without a passphrase.
func GenerateKeyPair(
	outputMode *io.OutputMode,
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

	return exec.NewCommandRunner(
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
	outputMode *io.OutputMode,
	projectDir string,
	dataDir string,
	force bool,
) error {
	instructorPublicKeyFile := filepath.Join(
		dataDir,
		InstructorPublicKeyFileName,
	)

	if instructorPublicKeyFileExists, err := filesystem.FileExists(
		instructorPublicKeyFile,
	); err != nil {
		return err
	} else if instructorPublicKeyFileExists {
		return fmt.Errorf(
			"instructor public key file exists at %s; "+
				"delete it first to regenerate keys",
			instructorPublicKeyFile,
		)
	}

	instructorPrivateKeyFile := filepath.Join(
		dataDir,
		InstructorPrivateKeyFileName,
	)

	instructorPrivateKeyFileExists, err := filesystem.FileExists(
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
