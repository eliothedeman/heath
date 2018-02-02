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

	y, err := NewBlock(nil, x, b)
	if err != nil {
		t.Error(err)
	}

	if y == nil {
		t.Fail()
	}

}
