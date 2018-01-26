package block

import (
	"crypto/ecdsa"
	"math/big"
)

func NewTransaction(priv *ecdsa.PrivateKey, payload []byte) (*Transaction, error) {
	sig, err := NewSignature(priv, payload)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		Signature: sig,
		Payload:   payload,
	}, nil
}

func (t *Transaction) Valid(keys []ecdsa.PublicKey) bool {
	ax := new(big.Int)
	bx := new(big.Int)
	ax.SetBytes(t.GetSignature().GetSignatureA())
	bx.SetBytes(t.GetSignature().GetSignatureB())

	// Iterate through keys to see
	for i := range keys {
		if ecdsa.Verify(&keys[i], hashPayload(t.GetPayload()), ax, bx) {
			return true
		}
	}
	return false
}

func signTransactions(p *ecdsa.PrivateKey, transactions []*Transaction) (*Signature, error) {
	h := hashTransactions(transactions)
	return NewSignature(p, h)
}

func hashTransactions(transactions []*Transaction) []byte {
	h := newHash()
	for _, t := range transactions {
		h.Write(t.GetPayload())
	}

	return h.Sum(nil)
}
