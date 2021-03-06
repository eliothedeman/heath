package db

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	"github.com/eliothedeman/heath/block"
	"github.com/eliothedeman/heath/block/test_util"
	"github.com/spf13/afero"
)

var testFS = afero.NewMemMapFs()

func newTestFile(name string) (io.ReadWriteSeeker, func() error) {
	f, _ := testFS.Create(name)
	return f, func() error {
		return f.Close()
	}
}

func newKey() *ecdsa.PrivateKey {
	k, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	return k
}

func TestDriverWrite(t *testing.T) {
	for name, df := range drivers {
		t.Run(fmt.Sprintf("Driver:%s", name), func(t *testing.T) {
			f, close := newTestFile(name)
			defer close()
			d := df(f)
			b := test_util.GenTestBlock(t, 3, 10, nil)

			err := d.Write(b)
			if err != nil {
				t.Error(err)
			}

			var x *block.Block
			xErr := EachBlock(d, func(y *block.Block) bool {
				if bytes.Equal(x.GetHash(), b.GetHash()) {
					x = y
					return false
				}
				return true
			})
			if xErr != nil {
				t.Error(xErr)
			}

			xx := x.GetTransactions()
			bb := b.GetTransactions()

			for i := range xx {
				if !bytes.Equal(xx[i].GetPayload(), bb[i].GetPayload()) {
					t.Error(*xx[i], *bb[i])
				}
			}

			close()
		})
	}
}

func TestDriversRead(t *testing.T) {
	for name, df := range drivers {
		t.Run(fmt.Sprintf("Driver:%s", name), func(t *testing.T) {
			f, close := newTestFile(name)
			defer close()
			d := df(f)

			var b *block.Block
			b = test_util.GenTestBlock(t, 1, 2, nil)
			for i := 0; i < 100; i++ {
				b = test_util.GenTestBlock(t, 1, 2, b)
				err := d.Write(b)
				if err != nil {
					t.Error(err)
				}
			}

			var count = 0
			out, err := d.StreamBlocks(context.Background())
			for _ = range out {
				count++
			}
			if count != 100 {
				t.Fail()
			}

			verr := <-err

			if verr != nil {
				t.Error(verr)
			}

			close()
		})
	}

}
