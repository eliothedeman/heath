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

func GenTestBlock(keys, transactions int, parent *Block) *Block {
	a, b := GenKeys(keys)
	x := GenTestTransactions(a, transactions)
	var hash *Hash
	if parent == nil {
		hash = parent.GetPetition().GetHash()
	}

	p := NewPetition(GenTestSignatures(a, x), x)
	y, _ := NewBlock(hash, p, x, b)

	return y
}
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

func GenTestTransaction(k *ecdsa.PrivateKey) *Transaction {
	t, _ := NewTransaction(k, randutil.Bytes(100), 0)
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
