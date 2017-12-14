package util

import (
	"crypto"
)

type SignerOpts struct{}

func (*SignerOpts) HashFunc() crypto.Hash {
	return crypto.SHA512
}

var SO = &SignerOpts{}
