package block

import "math/big"

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
