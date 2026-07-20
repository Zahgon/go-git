package sync

import (
	"bytes"
	"sync"
)

var (
	size = 32 * 1024

	byteSlice = sync.Pool{
		New: func() any {
			b := make([]byte, size)
			return &b
		},
	}
	bytesBuffer = sync.Pool{
		New: func() any {
			return bytes.NewBuffer(nil)
		},
	}
)

func GetByteSlice() *[]byte { _ = "STUB: not implemented"; return nil }

func PutByteSlice(buf *[]byte) { _ = "STUB: not implemented"; return }

func GetBytesBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func PutBytesBuffer(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }
