package transport

import (
	"context"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/storage"
)

type Commander interface {
	Command(ctx context.Context, cmd string, req packp.CommandArgs, resp packp.Decoder) error
}

type Transport interface {
	Handshake(ctx context.Context, req *Request) (Session, error)
}

type Session interface {
	Capabilities() *capability.List
	GetRemoteRefs(ctx context.Context, opts *GetRemoteRefsOptions) (*RemoteRefs, error)
	Fetch(ctx context.Context, st storage.Storer, req *FetchRequest) error
	Push(ctx context.Context, st storage.Storer, req *PushRequest) error
	Close() error
}

type GetRemoteRefsOptions struct {
	RefPrefixes []string
}

type RemoteRefs struct {
	References []*plumbing.Reference

	Unborn plumbing.ReferenceName
}

func NewRemoteRefs(refs []*plumbing.Reference) *RemoteRefs { _ = "STUB: not implemented"; return nil }
