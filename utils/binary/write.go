package binary

import (
	"io"
)

func Write(w io.Writer, data ...any) error { _ = "STUB: not implemented"; return nil }

func WriteVariableWidthInt(w io.Writer, n int64) error { _ = "STUB: not implemented"; return nil }

func WriteUint64(w io.Writer, value uint64) error { _ = "STUB: not implemented"; return nil }

func WriteUint32(w io.Writer, value uint32) error { _ = "STUB: not implemented"; return nil }

func WriteUint16(w io.Writer, value uint16) error { _ = "STUB: not implemented"; return nil }
