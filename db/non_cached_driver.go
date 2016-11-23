package db

import (
	"bytes"
	"context"
	"io"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/eliothedeman/heath/block"
)

const (
	defaultBufferSize = 1024 * 100
)

var (
	ErrBlockNotFound = errors.New("Block not found.")
)

func init() {
	RegisterDriver("non-cached", func(store io.ReadWriteSeeker, closeFunc func() error) Driver {

		return NewNonCachedDriver(store, closeFunc)
	})
}

// NonCachedDriver impliments the Driver interface with a single local file.
// As the name implies, this is a simple driver that does no caching.
type NonCachedDriver struct {
	rws   io.ReadWriteSeeker
	close func() error
	sync.Mutex
}

func NewNonCachedDriver(rws io.ReadWriteSeeker, closeFunc func() error) *NonCachedDriver {
	return &NonCachedDriver{
		rws:   rws,
		close: closeFunc,
	}
}

func (n *NonCachedDriver) WriteMulti(b []*block.Block) error {

	n.Lock()
	var err error
	for _, x := range b {
		err = n.writeNoLock(x)
		if err != nil {
			break
		}
	}
	n.Unlock()
	return err
}

func (n *NonCachedDriver) Write(b *block.Block) error {

	n.Lock()
	err := n.writeNoLock(b)
	n.Unlock()
	return err
}

func (n *NonCachedDriver) writeNoLock(b *block.Block) error {
	buff, err := proto.Marshal(b)
	if err != nil {
		return err
	}
	return writeWithPrefix(buff, n.rws)
}
func (n *NonCachedDriver) GetBlockByContentHash(hash []byte) (*block.Block, error) {
	ctx, cancel := context.WithCancel(context.Background())
	out, err := n.StreamBlocks(ctx)
	var b *block.Block
	for {
		select {
		case b = <-out:
			if bytes.Equal(b.GetSignature().ContentHash, hash) {
				cancel()
				return b, nil
			}
			break
		case e := <-err:
			cancel()
			return nil, e
		}
	}
}
func (n *NonCachedDriver) StreamBlocks(ctx context.Context) (<-chan *block.Block, <-chan error) {
	return n.SeekAndStream(time.Unix(0, 0), ctx)
}
func (n *NonCachedDriver) SeekAndStream(start time.Time, ctx context.Context) (<-chan *block.Block, <-chan error) {
	return n.SeekAndStreamUntil(start, time.Now(), ctx)
}

func (n *NonCachedDriver) SeekAndStreamUntil(t time.Time, until time.Time, ctx context.Context) (<-chan *block.Block, <-chan error) {
	out := make(chan *block.Block)
	errChan := make(chan error)

	go func() {
		n.Lock()
		var err error
		done := ctx.Done()
		defer func() {
			close(out)
			errChan <- err
			close(errChan)
			n.Unlock()
		}()
		buff := make([]byte, defaultBufferSize)
		_, err = n.rws.Seek(0, io.SeekStart)
		if err != nil {
			return
		}
		for {
			b := &block.Block{}
			err = readNextBlock(b, buff, n.rws)
			if err != nil {
				return
			}
			if b.GetSignature().GetTimestamp() > until.Unix() {
				return
			}

			select {
			case <-done:
				return
			case out <- b:
			}
		}

	}()

	return out, errChan
}
