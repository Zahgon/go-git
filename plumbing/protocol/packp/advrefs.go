package packp

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/protocol"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
)

type AdvRefs struct {
	Version protocol.Version

	Capabilities capability.List

	References []*plumbing.Reference

	Shallows []plumbing.Hash
}

func (a *AdvRefs) Head() (*plumbing.Reference, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *AdvRefs) ResolvedHead() (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AdvRefs) resolvedHeadFromSymref(head *plumbing.Reference) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AdvRefs) ResolvedReferences() ([]*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AdvRefs) symRefMap() (map[plumbing.ReferenceName]plumbing.ReferenceName, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AdvRefs) resolvedHeadFromHeuristic(head *plumbing.Reference) *plumbing.Reference {
	_ = "STUB: not implemented"
	return nil
}

func ResolveHeadFromHashHeuristic(head *plumbing.Reference, refs []*plumbing.Reference) *plumbing.Reference {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdvRefs) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (a *AdvRefs) supportSymrefs() bool { _ = "STUB: not implemented"; return false }
