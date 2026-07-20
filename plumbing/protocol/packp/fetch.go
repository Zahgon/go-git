package packp

import (
	"io"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
)

var maxSectionLines = 1 << 22

type MalformedResponseError struct {
	Reason string
}

func (e *MalformedResponseError) Error() string { _ = "STUB: not implemented"; return "" }

type FetchArgs struct {
	Wants []plumbing.Hash

	Haves []plumbing.Hash

	Done bool

	ThinPack bool

	NoProgress bool

	IncludeTag bool

	OFSDelta bool

	Shallows []plumbing.Hash

	Deepen int

	DeepenRelative bool

	DeepenSince time.Time

	DeepenNot []string

	Filter Filter

	WaitForDone bool
}

func (r *FetchArgs) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r *FetchArgs) Decode(rd io.Reader) error { _ = "STUB: not implemented"; return nil }

type Acknowledgments struct {
	ACKs []plumbing.Hash

	Ready bool
}

type ShallowInfo struct {
	Shallows []plumbing.Hash

	Unshallows []plumbing.Hash
}

type WantedRefs struct {
	Refs []*plumbing.Reference
}

type PackfileURIs struct {
	URIs []string
}

type FetchOutput struct {
	Acknowledgments *Acknowledgments

	ShallowInfo *ShallowInfo

	WantedRefs *WantedRefs

	PackfileURIs *PackfileURIs

	Packfile bool
}

func (r *FetchOutput) Decode(rd io.Reader) error { _ = "STUB: not implemented"; return nil }

func fetchSectionRank(header string) int { _ = "STUB: not implemented"; return 0 }

func (r *FetchOutput) decodeMetadataSection(rd io.Reader, decode func(io.Reader) (int, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *FetchOutput) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r *FetchOutput) decodeAcknowledgments(rd io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *FetchOutput) decodeShallowInfo(rd io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *FetchOutput) decodeWantedRefs(rd io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *FetchOutput) decodePackfileURIs(rd io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *FetchOutput) encodeAcknowledgments(w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *FetchOutput) encodeShallowInfo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r *FetchOutput) encodeWantedRefs(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r *FetchOutput) encodePackfileURIs(w io.Writer) error { _ = "STUB: not implemented"; return nil }
