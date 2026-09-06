package pktline

import (
	"io"

	"github.com/go-git/go-git/v6/utils/ioutil"
)

func Write(w io.Writer, p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func Writef(w io.Writer, format string, a ...any) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func Writeln(w io.Writer, s string) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func WriteString(w io.Writer, s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func WriteError(w io.Writer, e error) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func WriteFlush(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

func WriteDelim(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

func WriteResponseEnd(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

func Read(r io.Reader, p []byte) (l int, err error) { _ = "STUB: not implemented"; return 0, nil }

func ReadLine(r io.Reader) (l int, p []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func PeekLine(r ioutil.ReadPeeker) (l int, p []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func maskPackDataTrace(out bool, l int, data []byte) { _ = "STUB: not implemented"; return }
