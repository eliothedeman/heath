package block

import (
	"crypto/ecdsa"
	"testing"
)

func newKey() *ecdsa.PrivateKey {
	k, _ := GenerateKey()
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
