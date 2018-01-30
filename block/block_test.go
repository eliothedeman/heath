package block

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestBlockValid(t *testing.T) {
	a, b := GenKeys(10)
	x := GenTestTransactions(a, 10)
	p := NewPetition(GenTestSignatures(a, x), x)

	y, err := NewBlock(nil, p, x, b)
	if err != nil {
		t.Error(err)
	}

	if y == nil {
		t.Fail()
	}

}
