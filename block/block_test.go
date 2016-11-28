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
	a, b := genKeys(10)
	x := genTestTransactions(a, 89)
	p := NewPetition(genTestSignatures(a, x), x)

	y, err := NewBlock(nil, p, x, b)
	if err != nil {
		t.Error(err)
	}

	if y == nil {
		t.Fail()
	}

}
