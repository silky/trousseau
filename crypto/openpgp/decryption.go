package openpgp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "crypto/ecdsa"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	"code.google.com/p/go.crypto/openpgp"
	"code.google.com/p/go.crypto/openpgp/armor"
)

var keys openpgp.EntityList

func Decrypt(s, passphrase string) ([]byte, error) {
	if s == "" {
		return nil, nil
	}

	raw, err := armor.Decode(strings.NewReader(s))
	if err != nil {
		return nil, err
	}

	d, err := openpgp.ReadMessage(raw.Body, keys,
		func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
			kp := []byte(passphrase)

			if symmetric {
				return kp, nil
			}

			for _, k := range keys {
				err := k.PrivateKey.Decrypt(kp)
				if err == nil {
					return nil, nil
				}
			}

			return nil, fmt.Errorf("Unable to decrypt trousseau data store. " +
				"Invalid passphrase supplied.")
		},
		nil)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(d.UnverifiedBody)
	return bytes, err
}

func InitDecryption(keyRingPath, pass string) {
	f, err := os.Open(keyRingPath)
	if err != nil {
		log.Fatalf("unable to open gnupg keyring: %v", err)
	}
	defer f.Close()

	keys, err = openpgp.ReadKeyRing(f)
	if err != nil {
		log.Fatalf("unable to read from gnupg keyring: %v", err)
	}
}
