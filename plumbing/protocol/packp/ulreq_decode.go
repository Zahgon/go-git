package packp

import (
	"errors"
	"io"
)

var ErrDeepenMutuallyExclusive = errors.New("deepen and deepen-since (or deepen-not) cannot be used together")

func (req *UploadRequest) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }
