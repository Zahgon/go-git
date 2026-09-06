package packp

import (
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
)

type UploadRequest struct {
	Capabilities capability.List
	Wants        []plumbing.Hash
	Shallows     []plumbing.Hash
	Depth        DepthRequest
	Filter       Filter
}

type DepthRequest struct {
	Deepen int

	DeepenSince time.Time

	DeepenNot []string
}

func (d DepthRequest) IsZero() bool { _ = "STUB: not implemented"; return false }
