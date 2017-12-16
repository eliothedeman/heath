package keystore

import (
	"crypto/ecdsa"
	"errors"

	"github.com/eliothedeman/heath/block"
	"github.com/eliothedeman/heath/util"
)

var (
	// ErrNoPub is returned when a public key can't be found
	ErrNoPub = errors.New("Public key not found")
)

// Store provides storage and retreval of keys.
type Store interface {
	// Find attempts to find a public key given a signature.
	Find(sig *block.Signature) (*ecdsa.PublicKey, error)

	// Add A key to the store
	Add(key *ecdsa.PublicKey) error

	// Remove does what it sounds like
	Remove(key *ecdsa.PublicKey) error

	// Priv loads a private key
	Priv() *ecdsa.PrivateKey
}

// Memory is a store that stays in memory. Good for testing.
type Memory struct {
	pubs []*ecdsa.PublicKey
	my   *ecdsa.PrivateKey
}

// NewMemoryStore creates a new in memory Store
func NewMemoryStore(priv *ecdsa.PrivateKey) *Memory {
	return &Memory{
		my: priv,
	}
}

// Find a key by the given signature
func (m *Memory) Find(sig *block.Signature) (*ecdsa.PublicKey, error) {
	for _, p := range m.pubs {
		if ecdsa.Verify(p, sig.GetHash().GetContentHash(), sig.GetA(), sig.GetB()) {
			return p, nil
		}
	}

	return nil, ErrNoPub
}

// Add a key to the store
func (m *Memory) Add(key *ecdsa.PublicKey) error {
	m.pubs = append(m.pubs, key)
	return nil
}

func (m *Memory) Remove(key *ecdsa.PublicKey) error {
	for i, p := range m.pubs {
		if util.ComparePublicKeys(p, key) {
			m.pubs = append(m.pubs[:i], m.pubs[i+1:]...)
			return nil
		}
	}
	return ErrNoPub
}

func (m *Memory) Priv() *ecdsa.PrivateKey {
	return m.my
}
