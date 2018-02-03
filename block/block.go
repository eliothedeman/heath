package block

import (
	"crypto/ecdsa"
	"crypto/sha512"
	"hash"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrInvalidBlock = errors.New("Block Not Valid")
)

func newKey() *ecdsa.PrivateKey {
	k, _ := GenerateKey()
	return k
}

func NewHash() hash.Hash {
	return sha512.New()
}

func (b *Block) First() bool {
	return b.GetParent() == nil
}

func NewBlock(parent []byte, transactions []*Transaction) *Block {
	b := &Block{
		Timestamp:    time.Now().Unix(),
		Parent:       parent,
		Transactions: transactions,
	}

	h := NewHash()
	hashTransactions(h, transactions)
	b.Hash = h.Sum(nil)

	return b
}

func (b *Block) Valid(pubs []ecdsa.PublicKey) bool {

	// validate transactions
	for _, t := range b.GetTransactions() {
		if t == nil {
			return false
		}
		if !t.Valid(pubs) {
			return false
		}
	}

	return true
}
