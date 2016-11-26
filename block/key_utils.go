package block

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/golang/protobuf/proto"
)

func newCurve() elliptic.Curve {
	return elliptic.P521()
}

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(newCurve(), rand.Reader)
}

func MarshalKey(key *ecdsa.PrivateKey) ([]byte, error) {
	p := PublicKey{
		X: key.X.Bytes(),
		Y: key.Y.Bytes(),
	}
	k := PrivateKey{
		Public: &p,
		D:      key.D.Bytes(),
	}

	return proto.Marshal(&k)
}

func UnmarshalKey(buff []byte, key *ecdsa.PrivateKey) error {
	k := PrivateKey{}
	err := proto.Unmarshal(buff, &k)
	if err != nil {
		return err
	}

	key.X = new(big.Int)
	key.Y = new(big.Int)
	key.D = new(big.Int)
	key.Curve = newCurve()

	key.X.SetBytes(k.GetPublic().GetX())
	key.Y.SetBytes(k.GetPublic().GetY())
	key.D.SetBytes(k.GetD())

	return nil
}
