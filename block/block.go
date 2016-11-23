package block

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"math/big"
	"time"
)

func signPayload(priv *ecdsa.PrivateKey, payload []byte) (a, b []byte, hash []byte, err error) {
	h := sha512.Sum512(payload)
	hash = h[:]

	ax, bx, sErr := ecdsa.Sign(rand.Reader, priv, hash)
	a = ax.Bytes()
	b = bx.Bytes()
	err = sErr
	return
}

func now() *int64 {
	n := time.Now().UTC().Unix()
	return &n
}

func NewSignature(priv *ecdsa.PrivateKey, payload []byte) (*Signature, error) {
	a, b, hash, err := signPayload(priv, payload)
	if err != nil {
		return nil, err
	}

	return &Signature{
		SignatureA:  a,
		SignatureB:  b,
		ContentHash: hash,
		Timestamp:   now(),
	}, nil
}

func NewBlock(parent *Signature, priv *ecdsa.PrivateKey, payload []byte) (*Block, error) {
	sig, err := NewSignature(priv, payload)
	if err != nil {
		return nil, err
	}
	return &Block{
		Parent:    parent,
		Signature: sig,
		Payload:   payload,
	}, nil
}

func (b *Block) First() bool {
	return b.GetParent() == nil
}

func (b *Block) Valid(pub ecdsa.PublicKey) bool {
	ax := new(big.Int)
	bx := new(big.Int)
	ax.SetBytes(b.Signature.GetSignatureA())
	bx.SetBytes(b.Signature.GetSignatureB())
	return ecdsa.Verify(&pub, b.Signature.ContentHash, ax, bx)
}
