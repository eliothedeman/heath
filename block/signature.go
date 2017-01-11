package block

import (
	"crypto/ecdsa"
	"math/big"
)

func (s *Signature) GetA() *big.Int {
	i := new(big.Int)
	i.SetBytes(s.GetSignatureA())
	return i
}

func (s *Signature) GetB() *big.Int {
	i := new(big.Int)
	i.SetBytes(s.GetSignatureB())
	return i
}

func NewSignature(priv *ecdsa.PrivateKey, payload []byte) (*Signature, error) {
	a, b, hash, err := signPayload(priv, payload)
	if err != nil {
		return nil, err
	}

	return &Signature{
		Timestamp:  now(),
		SignatureA: a,
		SignatureB: b,
		Hash:       hash,
	}, nil
}
