package wire

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/eliothedeman/heath/util"
	"github.com/pkg/errors"
)

var (
	ErrNoKeyFound = errors.New("No key found")
)

func prepareOnetimeRsaKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func step1Client(priv *ecdsa.PrivateKey) (key *rsa.PrivateKey, hs Handshake1, err error) {
	key, err = prepareOnetimeRsaKey()
	if err != nil {
		return
	}
	hs.RsaPub, err = x509.MarshalPKIXPublicKey(key.PublicKey)
	if err != nil {
		return
	}
	hs.Signature, err = priv.Sign(rand.Reader, hs.RsaPub, util.SO)
	return
}

func step1Server(priv *ecdsa.PrivateKey, pubs []*ecdsa.PublicKey, hs *Handshake1) (key *rsa.PrivateKey, hs2 Handshake2, err error) {
	// validate the client connecting to us

}
