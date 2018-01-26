package block

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"hash"
	"time"

	"github.com/eliothedeman/randutil"
	"github.com/pkg/errors"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var (
	ErrInvalidBlock = errors.New("Block Not Valid")
)

func newKey() *ecdsa.PrivateKey {
	k, _ := GenerateKey()
	return k
}

func GenKeys(n int) ([]*ecdsa.PrivateKey, []ecdsa.PublicKey) {
	var priv []*ecdsa.PrivateKey
	var pub []ecdsa.PublicKey
	for i := 0; i < n; i++ {
		k := newKey()
		priv = append(priv, k)
		pub = append(pub, k.PublicKey)
	}

	return priv, pub
}

func newHash() hash.Hash {
	return sha512.New()
}

func hashPayload(payload []byte) []byte {
	return newHash().Sum(payload)
}

func signPayload(priv *ecdsa.PrivateKey, payload []byte) (a, b []byte, hash []byte, err error) {
	hash = hashPayload(payload)
	ax, bx, sErr := ecdsa.Sign(rand.Reader, priv, hash)
	a = ax.Bytes()
	b = bx.Bytes()
	err = sErr
	return
}

func now() *timestamp.Timestamp {
	t, _ := ptypes.TimestampProto(time.Now())
	return t
}

func (b *Block) First() bool {
	return b.GetParent() == nil
}

func NewBlock(parent []byte, petition *Petition, transactions []*Transaction, publicKeys []ecdsa.PublicKey) (*Block, error) {
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
		if t == nil {
			return false
		}
		if !t.Valid(pubs) {
			return false
		}
	}

	// validate petition
	p := b.GetPetition()
	if p == nil {
		return false
	}
	if !p.Valid(pubs) {
		return false
	}
	if !p.validateTransactions(b.GetTransactions()) {
		return false
	}

	return true
}

func GenTestBlock(keys, transactions int, parent *Block) *Block {
	a, b := GenKeys(keys)
	x := GenTestTransactions(a, transactions)
	var hash []byte
	if parent == nil {
		hash = parent.GetPetition().GetHash()
	}

	p := NewPetition(GenTestSignatures(a, x), x)
	y, _ := NewBlock(hash, p, x, b)

	return y
}

func GenTestTransaction(k *ecdsa.PrivateKey) *Transaction {
	t, _ := NewTransaction(k, randutil.Bytes(100))
	return t
}

func GenTestTransactions(keys []*ecdsa.PrivateKey, n int) []*Transaction {
	var t []*Transaction
	for i := 0; i < n; i++ {
		t = append(t, GenTestTransaction(keys[i%len(keys)]))
	}
	return t
}

func GenTestSignatures(keys []*ecdsa.PrivateKey, t []*Transaction) []*Signature {
	var s []*Signature

	for _, k := range keys {
		y, _ := signTransactions(k, t)
		s = append(s, y)
	}
	return s
}
