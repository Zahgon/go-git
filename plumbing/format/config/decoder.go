package config

import (
	"io"
)

type Decoder struct {
	io.Reader
}

func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

func (d *Decoder) Decode(config *Config) error { _ = "STUB: not implemented"; return nil }
