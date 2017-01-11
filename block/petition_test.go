package block

import "testing"

func TestPeitionValid(t *testing.T) {
	a, b := GenKeys(10)
	x := GenTestTransactions(a, 89)
	p := NewPetition(GenTestSignatures(a, x), x)

	if !p.Valid(b) {
		t.Fail()
	}

	if !p.validateTransactions(x) {
		t.Fail()
	}
}
