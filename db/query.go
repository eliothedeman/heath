package db

import (
	"context"

	"github.com/eliothedeman/heath/block"
)

// EachBlock iterates through all blocks the driver has access to.
// if "f" returns false, it will stop iterating
func EachBlock(d Driver, f func(b *block.Block) bool) error {
	c, cancel := context.WithCancel(context.Background())
	blocks, err := d.StreamBlocks(c)
	for b := range blocks {
		if !f(b) {
			cancel()
			return nil
		}
	}
	cancel()
	return <-err
}

// EachTransaction iterates through all transactions in all blocks
func EachTransaction(d Driver, f func(t *block.Transaction) bool) error {
	return EachBlock(d, func(b *block.Block) bool {
		for _, t := range b.Transactions {
			if !f(t) {
				return false
			}
		}
		return true
	})
}
