package test_util

import (
	"crypto/ecdsa"
	"testing"

	"github.com/eliothedeman/heath/block"
	"github.com/eliothedeman/heath/util"
	"github.com/eliothedeman/randutil"
)

func GenKeys(t *testing.T, n int) ([]*ecdsa.PrivateKey, []ecdsa.PublicKey) {
	var priv []*ecdsa.PrivateKey
	var pub []ecdsa.PublicKey
	for i := 0; i < n; i++ {
		k, err := util.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		priv = append(priv, k)
		pub = append(pub, k.PublicKey)
	}

	return priv, pub
}

func GenTestBlock(t *testing.T, keys, transactions int, parent *block.Block) *block.Block {
	a, _ := GenKeys(t, keys)
	x := GenTestTransactions(t, a, transactions)
	var hash []byte
	if parent != nil {
		hash = parent.Hash
	}

	return block.NewBlock(hash, x)
}

func GenTestTransaction(t *testing.T, k *ecdsa.PrivateKey) *block.Transaction {
	tx, err := block.NewTransaction(k, randutil.Bytes(100), block.Transaction_Raw)
	if err != nil {
		t.Fatal(err)
	}
	return tx
}

func GenTestTransactions(t *testing.T, keys []*ecdsa.PrivateKey, n int) []*block.Transaction {
	var tx []*block.Transaction
	for i := 0; i < n; i++ {
		tx = append(tx, GenTestTransaction(t, keys[i%len(keys)]))
	}
	return tx
}
