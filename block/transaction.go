package block

import (
	"crypto/ecdsa"
	"crypto/sha512"
	"hash"
	"math/big"
)

func NewTransaction(priv *ecdsa.PrivateKey, payload []byte, t Transaction_Type) (*Transaction, error) {
	h := NewHash()
	h.Write(payload)
	sig, err := NewSignature(priv, h)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		Signature: sig,
		Payload:   payload,
		Type:      t,
	}, nil
}

func (t *Transaction) Valid(keys []ecdsa.PublicKey) bool {
	ax := new(big.Int)
	bx := new(big.Int)
	ax.SetBytes(t.GetSignature().GetSignatureA())
	bx.SetBytes(t.GetSignature().GetSignatureB())

	b := make([]byte, sha512.Size)
	h := NewHash()
	h.Write(t.GetPayload())
	h.Sum(b[:0])
	// Iterate through keys to see
	for i := range keys {
		if ecdsa.Verify(&keys[i], b, ax, bx) {
			return true
		}
	}
	return false
}

func signTransactions(p *ecdsa.PrivateKey, transactions []*Transaction) (*Signature, error) {
	h := NewHash()
	hashTransactions(h, transactions)
	return NewSignature(p, h)
}

func hashTransactions(h hash.Hash, transactions []*Transaction) {
	for _, t := range transactions {
		h.Write(t.Payload)
	}
}
