package plumbing

import (
	"errors"
	"regexp"
)

const (
	refPrefix       = "refs/"
	refHeadPrefix   = refPrefix + "heads/"
	refTagPrefix    = refPrefix + "tags/"
	refRemotePrefix = refPrefix + "remotes/"
	refNotePrefix   = refPrefix + "notes/"
	symrefPrefix    = "ref: "
)

var RefRevParseRules = []string{
	"%s",
	"refs/%s",
	"refs/tags/%s",
	"refs/heads/%s",
	"refs/remotes/%s",
	"refs/remotes/%s/HEAD",
}

var (
	ErrReferenceNotFound = errors.New("reference not found")

	ErrInvalidReferenceName = errors.New("invalid reference name")
)

type ReferenceType int8

const (
	InvalidReference ReferenceType = 0

	HashReference ReferenceType = 1

	SymbolicReference ReferenceType = 2
)

func (r ReferenceType) String() string { _ = "STUB: not implemented"; return "" }

type ReferenceName string

func NewBranchReferenceName(name string) ReferenceName {
	_ = "STUB: not implemented"
	return *new(ReferenceName)
}

func NewNoteReferenceName(name string) ReferenceName {
	_ = "STUB: not implemented"
	return *new(ReferenceName)
}

func NewRemoteReferenceName(remote, name string) ReferenceName {
	_ = "STUB: not implemented"
	return *new(ReferenceName)
}

func NewRemoteHEADReferenceName(remote string) ReferenceName {
	_ = "STUB: not implemented"
	return *new(ReferenceName)
}

func NewTagReferenceName(name string) ReferenceName {
	_ = "STUB: not implemented"
	return *new(ReferenceName)
}

func (r ReferenceName) IsBranch() bool { _ = "STUB: not implemented"; return false }

func (r ReferenceName) IsNote() bool { _ = "STUB: not implemented"; return false }

func (r ReferenceName) IsRemote() bool { _ = "STUB: not implemented"; return false }

func (r ReferenceName) IsTag() bool { _ = "STUB: not implemented"; return false }

func (r ReferenceName) IsPeeled() bool { _ = "STUB: not implemented"; return false }

func (r ReferenceName) IsSafe() bool { _ = "STUB: not implemented"; return false }

func (r ReferenceName) String() string { _ = "STUB: not implemented"; return "" }

func (r ReferenceName) Short() string { _ = "STUB: not implemented"; return "" }

var ctrlSeqs = regexp.MustCompile(`[\000-\037\177]`)

func (r ReferenceName) Validate() error { _ = "STUB: not implemented"; return nil }

const (
	HEAD ReferenceName = "HEAD"

	Master ReferenceName = "refs/heads/master"

	Main ReferenceName = "refs/heads/main"

	Invalid ReferenceName = "refs/heads/.invalid"
)

type Reference struct {
	t      ReferenceType
	n      ReferenceName
	h      Hash
	target ReferenceName
}

func NewReferenceFromStrings(name, target string) *Reference { _ = "STUB: not implemented"; return nil }

func NewSymbolicReference(n, target ReferenceName) *Reference {
	_ = "STUB: not implemented"
	return nil
}

func NewHashReference(n ReferenceName, h Hash) *Reference { _ = "STUB: not implemented"; return nil }

func (r *Reference) Type() ReferenceType { _ = "STUB: not implemented"; return *new(ReferenceType) }

func (r *Reference) Name() ReferenceName { _ = "STUB: not implemented"; return *new(ReferenceName) }

func (r *Reference) Hash() Hash { _ = "STUB: not implemented"; return *new(Hash) }

func (r *Reference) Target() ReferenceName { _ = "STUB: not implemented"; return *new(ReferenceName) }

func (r *Reference) Strings() [2]string { _ = "STUB: not implemented"; return [2]string{} }

func (r *Reference) String() string { _ = "STUB: not implemented"; return "" }
