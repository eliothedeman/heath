package block

import (
	"crypto/ecdsa"
	"testing"

	"github.com/eliothedeman/randutil"
)

func genKeys(n int) ([]*ecdsa.PrivateKey, []ecdsa.PublicKey) {
	var priv []*ecdsa.PrivateKey
	var pub []ecdsa.PublicKey
	for i := 0; i < n; i++ {
		k := newKey()
		priv = append(priv, k)
		pub = append(pub, k.PublicKey)
	}

	return priv, pub
}

func genTestTransaction(k *ecdsa.PrivateKey) *Transaction {
	t, _ := NewTransaction(k, randutil.Bytes(100), 0)
	return t
}

func genTestTransactions(keys []*ecdsa.PrivateKey, n int) []*Transaction {
	var t []*Transaction
	for i := 0; i < n; i++ {
		t = append(t, genTestTransaction(keys[i%len(keys)]))
	}
	return t
}

func genTestSignatures(keys []*ecdsa.PrivateKey, t []*Transaction) []*Signature {
	var s []*Signature

	for _, k := range keys {
		y, _ := signTransactions(k, t)
		s = append(s, y)
	}
	return s
}

func TestPeitionValid(t *testing.T) {
	a, b := genKeys(10)
	x := genTestTransactions(a, 89)
	p := NewPetition(genTestSignatures(a, x), x)

	if !p.Valid(b) {
		t.Fail()
	}

	if !p.validateTransactions(x) {
		t.Fail()
	}
}
