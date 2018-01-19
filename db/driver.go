package db

import (
	"context"
	"time"

	"github.com/eliothedeman/heath/block"
)

var (
	drivers = map[string]DriverFactory{}
)

type DriverFactory func(url string, key *block.PrivateKey) Driver

func RegisterDriver(name string, df DriverFactory) {
	drivers[name] = df
}

// A Driver provides an interface to read and write blocks to/from persistance.
type Driver interface {
	StreamBlocks(ctx context.Context) (<-chan *block.Block, <-chan error)
	SeekAndStream(t time.Time, ctx context.Context) (<-chan *block.Block, <-chan error)
	SeekAndStreamUntil(start, until time.Time, ctx context.Context) (<-chan *block.Block, <-chan error)
	GetBlockByContentHash(hash []byte) (*block.Block, error)
	Write(b *block.Block) error
	WriteMulti(b []*block.Block) error
}
