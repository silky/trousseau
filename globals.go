package trousseau

import (
	"os"
	"path/filepath"
)

// Global data store file path
var gStorePath string = filepath.Join(os.Getenv("HOME"), STORE_FILENAME)

// Gnupg keyrings files
var gPubringFile string = filepath.Join(os.Getenv("HOME"), ".gnupg", "pubring.gpg")
var gSecringFile string = filepath.Join(os.Getenv("HOME"), ".gnupg", "secring.gpg")

// Ssh default identity file path
var gPrivateRsaKeyPath string = filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")

// Keyring manager service and username to use in order to
// retrieve trousseau main gpg key passphrase from system
// keyring
var gKeyringService string = os.Getenv(ENV_KEYRING_SERVICE_KEY)
var gKeyringUser string = os.Getenv(ENV_KEYRING_USER_KEY)
