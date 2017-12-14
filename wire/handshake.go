package wire

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"

	"github.com/pkg/errors"
)

var (
	ErrNoKeyFound = errors.New("No key found")
)

func prepareOnetimeRsaKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func step1Sender(priv *ecdsa.PrivateKey) (key *rsa.PrivateKey, signature []byte, err error) {
	k, err := prepareOnetimeRsaKey()
	if err != nil {
		return nil, nil, err
	}
}
