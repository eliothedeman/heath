package wire

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/eliothedeman/heath/block"

	"github.com/eliothedeman/heath/keystore"
	"github.com/pkg/errors"
)

var (
	ErrNoKeyFound = errors.New("No key found")
)

func prepareOnetimeRsaKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func step1Client(store keystore.Store) (key *rsa.PrivateKey, hs Handshake1, err error) {
	key, err = prepareOnetimeRsaKey()
	if err != nil {
		return
	}
	hs.RsaPub, err = x509.MarshalPKIXPublicKey(key.PublicKey)
	if err != nil {
		return
	}

	hs.Signature, err = block.NewSignature(store.Priv(), hs.RsaPub)
	return
}

func step1Server(store keystore.Store, hs *Handshake1) (key *rsa.PrivateKey, pub *rsa.PublicKey, hs2 Handshake2, err error) {
	// validate the client connecting to us
	_, err = store.Find(hs.Signature)
	if err != nil {
		return
	}

	var i interface{}
	i, err = x509.ParsePKIXPublicKey(hs.RsaPub)
	if err != nil {
		return
	}
	pub = i.(*rsa.PublicKey)

	key, err = prepareOnetimeRsaKey()
	if err != nil {
		return
	}

	return
}
