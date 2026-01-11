package ssh

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/exec"
	"github.com/TypingHare/course-sync/internal/infra/fs"
)

// GenerateKeyPair generates a new SSH key pair. This function ensures that the
// parent directory for the private key file exists before generating the
// pair. Then, it uses the ssh-keygen command to create an ed25519 key pair
// without a passphrase.
func GenerateKeyPair(appCtx *app.Context, privateKeyFile string) error {
	exec.ShellEnsureDir(appCtx, filepath.Dir(privateKeyFile))

	commandTask := exec.NewCommandTask(
		appCtx,
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
	)

	_, err := commandTask.Start()

	return err
}

// GenerateMasterKeyPair generates a new master SSH key pair and saves it to the
// application directory.
func GenerateMasterKeyPair(appCtx *app.Context, force bool) error {
	masterPrivateKeyFile := filepath.Join(
		appCtx.AppDataDir,
		app.MASTER_PRIVATE_KEY_FILE_NAME,
	)

	masterPrivateKeyFileExists, err := fs.FileExists(masterPrivateKeyFile)
	if err != nil {
		return err
	}

	if masterPrivateKeyFileExists && !force {
		return fmt.Errorf(
			"master private key file already exists at %s",
			masterPrivateKeyFile,
		)
	}

	return GenerateKeyPair(appCtx, masterPrivateKeyFile)
}
