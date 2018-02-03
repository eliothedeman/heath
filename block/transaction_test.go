package block_test

import (
	"crypto/ecdsa"
	"testing"

	"github.com/eliothedeman/heath/block"
	"github.com/eliothedeman/heath/util"
)

func TestTransactionValid(t *testing.T) {
	msg := "hello world"
	k, _ := util.GenerateKey()
	tx, err := block.NewTransaction(k, []byte(msg), block.Transaction_Raw)
	if err != nil {
		t.Error(tx)
	}
	if !tx.Valid([]ecdsa.PublicKey{k.PublicKey}) {
		t.Error("Transaction not valid")
	}

	nk, _ := util.GenerateKey()
	if tx.Valid([]ecdsa.PublicKey{nk.PublicKey}) {
		t.Error("Transaction should not be valid")
	}

	got := tx.GetPayload()

	if string(got) != msg {
		t.Errorf("wanted %s got %s", msg, string(got))
	}

}
