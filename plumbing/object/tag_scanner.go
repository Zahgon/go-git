package object

import (
	"bufio"
	"bytes"
)

type tagScanner struct {
	r      *bufio.Reader
	t      *Tag
	msgbuf bytes.Buffer

	pending    []byte
	pendingErr error

	sawObject, sawType, sawName, sawTagger bool
}

type tagState func(*tagScanner) (tagState, error)

func (s *tagScanner) readLine() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *tagScanner) pushBack(line []byte, err error) { _ = "STUB: not implemented"; return }

func scanTagObject(s *tagScanner) (tagState, error) {
	_ = "STUB: not implemented"
	return *new(tagState), nil
}

func scanTagType(s *tagScanner) (tagState, error) {
	_ = "STUB: not implemented"
	return *new(tagState), nil
}

func scanTagName(s *tagScanner) (tagState, error) {
	_ = "STUB: not implemented"
	return *new(tagState), nil
}

func scanTagTagger(s *tagScanner) (tagState, error) {
	_ = "STUB: not implemented"
	return *new(tagState), nil
}

func scanTagHeaders(s *tagScanner) (tagState, error) {
	_ = "STUB: not implemented"
	return *new(tagState), nil
}

func scanTagPgp256Cont(s *tagScanner) (tagState, error) {
	_ = "STUB: not implemented"
	return *new(tagState), nil
}

func scanTagMessage(s *tagScanner) (tagState, error) {
	_ = "STUB: not implemented"
	return *new(tagState), nil
}
