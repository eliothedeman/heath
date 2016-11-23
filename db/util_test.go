package db

import (
	"bytes"
	"testing"
)

func TestReadWriteFull(t *testing.T) {
	b := bytes.NewBuffer(nil)
	msg := "hello wrold"
	err := writeWithPrefix([]byte(msg), b)
	if err != nil {
		t.Error(err)
	}

	out := make([]byte, 1024)
	var n int
	n, err = readWithPrefix(out, b)
	if err != nil {
		t.Error(err)
	}

	if string(out[:n]) != msg {
		t.Error(string(out[:n]))
	}

}
