package db

import (
	"encoding/binary"
	"io"

	"github.com/eliothedeman/heath/block"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var (
	enc            = binary.LittleEndian
	ErrBuffToSmall = errors.New("Buffer too small for requested opperation.")
)

func writeFull(buff []byte, w io.Writer) error {
	var n int
	var i int
	var err error

	for i < len(buff) {
		n, err = w.Write(buff)
		if err != nil {
			return err
		}
		i += n
	}

	return nil
}

func getPrefix(buff, prefixBuff []byte) {
	enc.PutUint64(prefixBuff, uint64(len(buff)))
}

func readNextPrefix(r io.Reader) (int, error) {
	prefixBuff := make([]byte, 8)
	err := readFull(prefixBuff, r)
	return int(enc.Uint64(prefixBuff)), err
}

func writeWithPrefix(buff []byte, w io.Writer) error {
	prefixBuff := make([]byte, 8)
	getPrefix(buff, prefixBuff)
	err := writeFull(prefixBuff, w)
	if err != nil {
		return err
	}

	return writeFull(buff, w)
}

func readWithPrefix(buff []byte, r io.Reader) (int, error) {
	size, err := readNextPrefix(r)
	if err != nil {
		return -1, err
	}
	if size > len(buff) {
		return -1, errors.Wrapf(ErrBuffToSmall, "Needed at least %d got %d", size, len(buff))
	}

	return size, readFull(buff[:size], r)
}

func readFull(buff []byte, r io.Reader) error {
	var n int
	var i int
	var err error
	for i < len(buff) {
		n, err = r.Read(buff)
		if err != nil {
			return err
		}
		i += n
	}

	return nil
}

func readNextBlock(b *block.Block, buff []byte, r io.Reader) error {
	n, err := readWithPrefix(buff, r)
	if err != nil {
		return err
	}

	return proto.Unmarshal(buff[:n], b)
}
