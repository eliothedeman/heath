package block

import (
	"crypto/ecdsa"
	"crypto/rand"
	"hash"
	"math/big"
	"time"
)

var (
	reader = rand.Reader
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

func NewSignature(priv *ecdsa.PrivateKey, h hash.Hash) (*Signature, error) {
	s := new(Signature)
	s.Timestamp = time.Now().Unix()
	err := sign(priv, h, s)
	if err != nil {
		return nil, err
	}
	return s, err
}

func sign(priv *ecdsa.PrivateKey, h hash.Hash, s *Signature) error {
	s.Hash = h.Sum(s.Hash)
	ax, bx, sErr := ecdsa.Sign(reader, priv, s.Hash)
	s.SignatureA = ax.Bytes()
	s.SignatureB = bx.Bytes()
	return sErr
}
