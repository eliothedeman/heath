package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func NewCurve() elliptic.Curve {
	return elliptic.P521()
}

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(NewCurve(), rand.Reader)
}

func ComparePublicKeys(a, b *ecdsa.PublicKey) bool {
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}
