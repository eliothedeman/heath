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
	RegisterDriver("non-cached", func(backing io.ReadWriteSeeker) Driver {
		return NewNonCachedDriver(backing)
	})
}

// NonCachedDriver impliments the Driver interface with a single local file.
// As the name implies, this is a simple driver that does no caching.
type NonCachedDriver struct {
	rws io.ReadWriteSeeker
	sync.Mutex
}

func NewNonCachedDriver(rws io.ReadWriteSeeker) *NonCachedDriver {
	return &NonCachedDriver{
		rws: rws,
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
	n.rws.Seek(0, io.SeekEnd)
	return writeWithPrefix(buff, n.rws)
}
func (n *NonCachedDriver) GetBlockByContentHash(hash []byte) (*block.Block, error) {
	ctx, cancel := context.WithCancel(context.Background())
	out, err := n.StreamBlocks(ctx)
	var b *block.Block
	for {
		select {
		case b = <-out:
			if bytes.Equal(b.GetPetition().GetHash(), hash) {
				cancel()
				return b, <-err
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

	// create output channels
	out := make(chan *block.Block)
	errChan := make(chan error)

	go func() {
		n.Lock()

		// this will be used to signal a cancle or timeout
		done := ctx.Done()

		// send an error on the err channel reguardless.
		var err error
		defer func() {
			close(out)
			if err == io.EOF {
				err = nil
			}
			select {
			case errChan <- err:

			case <-done:

			}
			close(errChan)
			n.Unlock()
		}()

		buff := make([]byte, defaultBufferSize)

		// allways seek to the beginning to start the read.
		_, err = n.rws.Seek(0, io.SeekStart)
		if err != nil {
			return
		}
		for {

			// TODO maybe think about recycling these blocks.
			b := &block.Block{}
			err = readNextBlock(b, buff, n.rws)
			if err != nil {
				return
			}

			// Ensure that the timestamp is before the deadline.
			if b.GetTimestamp().Seconds > until.Unix() {
				return
			}

			select {
			case <-done:
				err = ctx.Err()
				return
			case out <- b:
			}
		}

	}()

	return out, errChan
}
