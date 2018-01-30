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

	h := block.NewHash()
	h.Write([]byte("Hello world"))

	x, err := block.NewSignature(a, h)
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
	h2 := block.NewHash()
	h2.Write([]byte("whats up"))
	x, err = block.NewSignature(a, h2)
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
