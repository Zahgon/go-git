package config

import (
	"errors"

	"github.com/go-git/go-git/v6/plumbing"
	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

var (
	errBranchEmptyName     = errors.New("branch config: empty name")
	errBranchInvalidRebase = errors.New("branch config: rebase must be one of 'true' or 'interactive'")
)

type Branch struct {
	Name string

	Remote string

	Merge plumbing.ReferenceName

	Rebase string

	Description string

	raw *format.Subsection
}

func (b *Branch) Validate() error { _ = "STUB: not implemented"; return nil }

func (b *Branch) marshal() *format.Subsection { _ = "STUB: not implemented"; return nil }

func quoteDescription(desc string) string { _ = "STUB: not implemented"; return "" }

func (b *Branch) unmarshal(s *format.Subsection) error { _ = "STUB: not implemented"; return nil }

func unquoteDescription(desc string) string { _ = "STUB: not implemented"; return "" }
