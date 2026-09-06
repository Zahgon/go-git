package sideband

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing/format/pktline"
)

var ErrMaxPackedExceeded = errors.New("max. packed size exceeded")

type Progress interface {
	io.Writer
}

type Demuxer struct {
	t Type
	r io.Reader
	s *pktline.Scanner

	max     int
	pending []byte

	Progress Progress
}

func NewDemuxer(t Type, r io.Reader) *Demuxer { _ = "STUB: not implemented"; return nil }

func (d *Demuxer) Read(b []byte) (read int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d *Demuxer) doRead(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *Demuxer) nextPackData() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Demuxer) getPending() (b []byte) { _ = "STUB: not implemented"; return nil }
