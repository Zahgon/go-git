package config

import (
	"errors"

	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

var errURLEmptyInsteadOf = errors.New("url config: empty insteadOf")

type URL struct {
	Name string

	InsteadOfs []string

	raw *format.Subsection
}

func (u *URL) Validate() error { _ = "STUB: not implemented"; return nil }

const (
	insteadOfKey = "insteadOf"
)

func (u *URL) unmarshal(s *format.Subsection) error { _ = "STUB: not implemented"; return nil }

func (u *URL) marshal() *format.Subsection { _ = "STUB: not implemented"; return nil }

func applyLongestInsteadOfMatch(remoteURL string, urls []*URL) (rewrittenURL string, matched bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (u *URL) ApplyInsteadOf(url string) string { _ = "STUB: not implemented"; return "" }
