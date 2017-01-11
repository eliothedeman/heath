package block

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"hash"
	"time"

	"github.com/pkg/errors"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var (
	ErrInvalidBlock = errors.New("Block Not Valid")
)

func GenTestBlock(keys, transactions int, parent *Block) *Block {
	a, b := GenKeys(keys)
	x := GenTestTransactions(a, transactions)

	p := NewPetition(GenTestSignatures(a, x), x)
	y, _ := NewBlock(parent.Petition.GetHash(), p, x, b)

	return y
}

func newHash() hash.Hash {
	return sha512.New()
}

func hashPayload(payload []byte) *Hash {
	return &Hash{
		ContentHash: newHash().Sum(payload),
	}
}

func signPayload(priv *ecdsa.PrivateKey, payload []byte) (a, b []byte, hash *Hash, err error) {
	hash = hashPayload(payload)
	ax, bx, sErr := ecdsa.Sign(rand.Reader, priv, hash.ContentHash)
	a = ax.Bytes()
	b = bx.Bytes()
	err = sErr
	return
}

func now() *timestamp.Timestamp {
	t, _ := ptypes.TimestampProto(time.Now())
	return t
}

func NewSignature(priv *ecdsa.PrivateKey, payload []byte) (*Signature, error) {
	a, b, hash, err := signPayload(priv, payload)
	if err != nil {
		return nil, err
	}

	return &Signature{
		SignatureA: a,
		SignatureB: b,
		Hash:       hash,
	}, nil
}

func (b *Block) First() bool {
	return b.GetParent() == nil
}

func NewBlock(parent *Hash, petition *Petition, transactions []*Transaction, publicKeys []ecdsa.PublicKey) (*Block, error) {
	b := Block{
		Timestamp:    now(),
		Parent:       parent,
		Petition:     petition,
		Transactions: transactions,
	}

	if !b.Valid(publicKeys) {
		return nil, ErrInvalidBlock
	}

	return &b, nil
}

func (b *Block) Valid(pubs []ecdsa.PublicKey) bool {

	// validate transactions
	for _, t := range b.GetTransactions() {
		if !t.Valid(pubs) {
			return false
		}
	}

	// validate petition
	if !b.GetPetition().Valid(pubs) {
		return false
	}
	if !b.GetPetition().validateTransactions(b.GetTransactions()) {
		return false
	}

	return true
}
