package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/pktline"
)

const (
	ok = "ok"
)

type UnpackStatusErr struct {
	Status string
}

func (e UnpackStatusErr) Error() string { _ = "STUB: not implemented"; return "" }

type CommandStatusErr struct {
	ReferenceName plumbing.ReferenceName
	Status        string
}

func (e CommandStatusErr) Error() string { _ = "STUB: not implemented"; return "" }

type ReportStatus struct {
	UnpackStatus    string
	CommandStatuses []*CommandStatus
}

func (s *ReportStatus) Error() error { _ = "STUB: not implemented"; return nil }

func (s *ReportStatus) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (s *ReportStatus) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (s *ReportStatus) scanFirstLine(sc *pktline.Scanner) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ReportStatus) decodeReportStatus(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *ReportStatus) decodeCommandStatus(b []byte) error { _ = "STUB: not implemented"; return nil }

type CommandStatus struct {
	ReferenceName plumbing.ReferenceName
	Status        string
}

func (s *CommandStatus) Error() error { _ = "STUB: not implemented"; return nil }

func (s *CommandStatus) encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }
