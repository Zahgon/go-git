package pktline

import "sync"

var pktBuffer = sync.Pool{
	New: func() any {
		var b [MaxSize]byte
		return &b
	},
}

func GetBuffer() *[MaxSize]byte { _ = "STUB: not implemented"; return nil }

func PutBuffer(buf *[MaxSize]byte) { _ = "STUB: not implemented"; return }
