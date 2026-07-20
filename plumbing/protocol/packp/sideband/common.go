package sideband

type Type int8

const (
	Sideband Type = iota

	Sideband64k Type = iota

	MaxPackedSize = 1000

	MaxPackedSize64k = 65520
)

type Channel byte

func (ch Channel) WithPayload(payload []byte) []byte { _ = "STUB: not implemented"; return nil }

const (
	PackData Channel = 1

	ProgressMessage Channel = 2

	ErrorMessage Channel = 3
)
