package sync

import (
	"bufio"
	"io"
	"sync"
)

var bufioReader = sync.Pool{
	New: func() any {
		return bufio.NewReader(nil)
	},
}

func GetBufioReader(reader io.Reader) *bufio.Reader { _ = "STUB: not implemented"; return nil }

func PutBufioReader(reader *bufio.Reader) { _ = "STUB: not implemented"; return }
