package packp

import "io"

type Encoder interface {
	Encode(w io.Writer) error
}

type Decoder interface {
	Decode(r io.Reader) error
}
