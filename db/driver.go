package db

import (
	"context"
	"io"

	"github.com/eliothedeman/heath/block"
)

var (
	drivers = map[string]DriverFactory{}
)

type DriverFactory func(io.ReadWriteSeeker) Driver

func RegisterDriver(name string, df DriverFactory) {
	drivers[name] = df
}

func Open(kind string, backing io.ReadWriteSeeker) Driver {
	return NewNonCachedDriver(backing)
}

// A Driver provides an interface to read and write blocks to/from persistance.
type Driver interface {
	StreamBlocks(ctx context.Context) (<-chan *block.Block, <-chan error)
	Write(b *block.Block) error
}
