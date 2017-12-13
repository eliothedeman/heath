package block

import (
	"crypto/ecdsa"
	"testing"
)

func TestTransactionValid(t *testing.T) {
	msg := "hello world"
	k := newKey()
	tx, err := NewTransaction(k, []byte(msg))
	if err != nil {
		t.Error(tx)
	}
	if !tx.Valid([]ecdsa.PublicKey{k.PublicKey}) {
		t.Error("Transaction not valid")
	}

	nk := newKey()
	if tx.Valid([]ecdsa.PublicKey{nk.PublicKey}) {
		t.Error("Transaction should not be valid")
	}

	if string(tx.GetPayload()) != msg {
		t.Errorf("wanted %s got %s", msg, string(tx.GetPayload()))
	}

}
