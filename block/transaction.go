package block

import (
	"crypto/ecdsa"
	"crypto/sha512"
	"math/big"
)

func NewTransaction(priv *ecdsa.PrivateKey, payload []byte, payloadType int64) (*Transaction, error) {
	sig, err := NewSignature(priv, payload)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		Signature:   sig,
		Payload:     payload,
		PayloadType: &payloadType,
	}, nil
}

func (t *Transaction) Valid(keys []ecdsa.PublicKey) bool {
	ax := new(big.Int)
	bx := new(big.Int)
	ax.SetBytes(t.GetSignature().GetSignatureA())
	bx.SetBytes(t.GetSignature().GetSignatureB())

	// Iterate through keys to see
	for i := range keys {
		if ecdsa.Verify(&keys[i], hashPayload(t.GetPayload()).GetContentHash(), ax, bx) {
			return true
		}
	}
	return false

}

func hashTransactions(transactions []*Transaction) []byte {
	h := sha512.New()
	for _, t := range transactions {
		h.Write(t.GetPayload())
	}

	return h.Sum(nil)
}
