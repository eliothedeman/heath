package keystore

import (
	"crypto/ecdsa"
	"testing"

	"github.com/eliothedeman/heath/block"

	"github.com/eliothedeman/heath/util"
)

func genKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()
	a, err := util.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func testStore(t *testing.T, s Store) {
	t.Helper()

	a := genKey(t)
	err := s.Add(&a.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	x, err := block.NewSignature(a, []byte("Hello world"))
	if err != nil {
		t.Fatal(err)
	}

	k, err := s.Find(x)
	if err != nil {
		t.Fatal(err)
	}

	if !util.ComparePublicKeys(&a.PublicKey, k) {
		t.Fatal("Wrong key stored")
	}

	err = s.Remove(k)
	if err != nil {
		t.Fatal(err)
	}
	x, err = block.NewSignature(a, []byte("Whats up"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.Find(x)
	if err != ErrNoPub {
		t.Fatal("Public key should be gone")
	}
}

func TestMemoryStore(t *testing.T) {
	p, err := util.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	m := NewMemoryStore(p)
	testStore(t, m)
}
