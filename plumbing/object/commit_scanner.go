package object

import (
	"bufio"
	"bytes"

	"github.com/go-git/go-git/v6/plumbing"
)

type commitScanner struct {
	r      *bufio.Reader
	c      *Commit
	msgbuf bytes.Buffer

	pending    []byte
	pendingErr error

	sawTree, sawAuthor, sawCommitter bool
	sawEncoding                      bool

	extra *ExtraHeader
}

type commitState func(*commitScanner) (commitState, error)

func (s *commitScanner) readLine() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *commitScanner) pushBack(line []byte, err error) { _ = "STUB: not implemented"; return }

func scanTree(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func scanParents(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func scanAuthor(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func scanCommitter(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func scanHeaders(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func scanPgpCont(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func scanPgp256Cont(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func continuationCont(s *commitScanner, dst *string, self commitState) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func scanExtraCont(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func (s *commitScanner) finaliseExtra() { _ = "STUB: not implemented"; return }

func scanMessage(s *commitScanner) (commitState, error) {
	_ = "STUB: not implemented"
	return *new(commitState), nil
}

func isBlankLine(line []byte) bool { _ = "STUB: not implemented"; return false }

func splitHeader(line []byte) (string, []byte) { _ = "STUB: not implemented"; return "", nil }

func parseObjectIDHex(data []byte, malformedErr error, header string) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}
