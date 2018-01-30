package block

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha512"
)

func NewPetition(signatures []*Signature, transactions []*Transaction) *Petition {
	h := newHash()
	b := make([]byte, sha512.Size)
	hashTransactions(h, transactions)
	b = h.Sum(b[:0])

	return &Petition{
		Hash:       b,
		Signatures: signatures,
	}
}

func (p *Petition) validateTransactions(transactions []*Transaction) bool {
	// calculate hash of transactions
	h := newHash()
	hashTransactions(h, transactions)
	b := make([]byte, sha512.Size)

	if !bytes.Equal(p.GetHash(), h.Sum(b[:0])) {
		return false
	}

	return true
}

// Valid ensures that every signature in the petition is a vaild signature of
// the petition.
func (p *Petition) Valid(keys []ecdsa.PublicKey) bool {

	for _, s := range p.GetSignatures() {
		valid := false
		for i := range keys {
			if ecdsa.Verify(&keys[i], s.GetHash(), s.GetA(), s.GetB()) {
				valid = true
				break
			}
		}
		if !valid {
			return false
		}
	}

	return true
}
