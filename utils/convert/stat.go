package convert

import "io"

type Stat struct {
	NUL, LoneCR, LoneLF, CRLF uint
	Printable, NonPrintable   uint
}

func (s Stat) IsBinary() bool { _ = "STUB: not implemented"; return false }

func GetStat(r io.Reader) (stat Stat, err error) { _ = "STUB: not implemented"; return *new(Stat), nil }
