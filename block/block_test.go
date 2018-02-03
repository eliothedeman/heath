package block_test

import (
	"log"
	"testing"

	"github.com/eliothedeman/heath/block"
	. "github.com/eliothedeman/heath/block/test_util"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestBlockValid(t *testing.T) {
	a, _ := GenKeys(t, 10)
	x := GenTestTransactions(t, a, 10)

	y := block.NewBlock(nil, x)

	if y == nil {
		t.Fail()
	}

}
