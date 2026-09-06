package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
)

func EncodeListV2(w io.Writer, l *capability.List) error { _ = "STUB: not implemented"; return nil }

func DecodeListV2(r io.Reader, l *capability.List) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
