package block

import (
	"bytes"
	"crypto/ecdsa"
)

func NewPetition(signatures []*Signature, transactions []*Transaction) *Petition {
	hashedContent := hashTransactions(transactions)

	return &Petition{
		Hash:       hashedContent,
		Signatures: signatures,
	}
}

func (p *Petition) validateTransactions(transactions []*Transaction) bool {
	// calculate hash of transactions
	h := newHash()
	for _, t := range transactions {
		h.Write(t.GetPayload())
	}

	if !bytes.Equal(p.GetHash(), h.Sum(nil)) {
		return false
	}

	return true
}

// Valid ensures that every signature in the petition is a vaild signature of
// the petition.
func (p *Petition) Valid(keys []ecdsa.PublicKey) bool {

	for _, s := range p.GetSignatures() {
		valid := false
		for i, _ := range keys {
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
