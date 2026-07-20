package config

import (
	"errors"

	"github.com/go-git/go-git/v6/plumbing"
)

const (
	refSpecWildcard  = "*"
	refSpecForce     = "+"
	refSpecSeparator = ":"
)

var (
	ErrRefSpecMalformedSeparator = errors.New("malformed refspec, separators are wrong")

	ErrRefSpecMalformedWildcard = errors.New("malformed refspec, mismatched number of wildcards")
)

type RefSpec string

func (s RefSpec) Validate() error { _ = "STUB: not implemented"; return nil }

func (s RefSpec) IsForceUpdate() bool { _ = "STUB: not implemented"; return false }

func (s RefSpec) IsDelete() bool { _ = "STUB: not implemented"; return false }

func (s RefSpec) IsExactSHA1() bool { _ = "STUB: not implemented"; return false }

func (s RefSpec) Src() string { _ = "STUB: not implemented"; return "" }

func (s RefSpec) Match(n plumbing.ReferenceName) bool { _ = "STUB: not implemented"; return false }

func (s RefSpec) IsWildcard() bool { _ = "STUB: not implemented"; return false }

func (s RefSpec) matchExact(n plumbing.ReferenceName) bool { _ = "STUB: not implemented"; return false }

func (s RefSpec) matchGlob(n plumbing.ReferenceName) bool { _ = "STUB: not implemented"; return false }

func (s RefSpec) Dst(n plumbing.ReferenceName) plumbing.ReferenceName {
	_ = "STUB: not implemented"
	return *new(plumbing.ReferenceName)
}

func (s RefSpec) Reverse() RefSpec { _ = "STUB: not implemented"; return *new(RefSpec) }

func (s RefSpec) String() string { _ = "STUB: not implemented"; return "" }

func MatchAny(l []RefSpec, n plumbing.ReferenceName) bool { _ = "STUB: not implemented"; return false }
