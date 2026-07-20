package pktline

import (
	"errors"
)

const (
	Err = iota - 1

	Flush

	Delim

	ResponseEnd
)

const (
	MaxPayloadSize = MaxSize - LenSize

	MaxSize = 65520

	LenSize = 4
)

var (
	ErrPayloadTooLong = errors.New("payload is too long")

	ErrInvalidPktLen = errors.New("invalid pkt-len found")
)

var (
	flushPkt = []byte{'0', '0', '0', '0'}

	delimPkt = []byte{'0', '0', '0', '1'}

	responseEndPkt = []byte{'0', '0', '0', '2'}

	emptyPkt = []byte{'0', '0', '0', '4'}
)
