package block

import (
	"crypto/ecdsa"
	"crypto/sha512"
	"hash"
	"log"
	"math/big"
)

func NewTransaction(priv *ecdsa.PrivateKey, payload isTransaction_Payload) (*Transaction, error) {
	h := newHash()
	hashPayload(h, payload)
	sig, err := NewSignature(priv, h)
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

	b := make([]byte, sha512.Size)
	h := newHash()
	hashPayload(h, t.GetPayload())
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
	h := newHash()
	hashTransactions(h, transactions)
	return NewSignature(p, h)
}
func hashPayload(h hash.Hash, p isTransaction_Payload) {
	switch p := p.(type) {
	case *Transaction_Raw:
		h.Write(p.Raw)
	case *Transaction_NewPublicKey:
		k := p.NewPublicKey
		h.Write(k.GetX())
		h.Write(k.GetY())
	case nil:
	default:
		log.Fatalf("Unknown type %t", p)

	}
}

func hashTransactions(h hash.Hash, transactions []*Transaction) {
	for _, t := range transactions {
		hashPayload(h, t.GetPayload())
	}
}
