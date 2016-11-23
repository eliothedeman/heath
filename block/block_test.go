package block

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func newKey() *ecdsa.PrivateKey {
	k, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	return k
}

func TestBlockValid(t *testing.T) {
	k := newKey()
	b, err := NewBlock(nil, k, []byte("hello world"))
	if err != nil {
		t.Error(err)
	}

	if !b.First() {
		t.Fail()
	}

	if !b.Valid(k.PublicKey) {
		t.Fail()
	}

	if b.Valid(newKey().PublicKey) {
		t.Fail()
	}
}
