package pktline

import (
	"io"
)

type Scanner struct {
	r   io.Reader
	err error
	buf [MaxSize]byte
	n   int
}

func NewScanner(r io.Reader) *Scanner { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Err() error { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Scan() bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Text() string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) Len() int { _ = "STUB: not implemented"; return 0 }
