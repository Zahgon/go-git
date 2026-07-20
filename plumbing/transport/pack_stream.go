package transport

import (
	"bufio"
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/storage"
)

type StreamSession struct {
	conn    Conn
	r       *bufio.Reader
	w       io.WriteCloser
	svc     string
	version protocol.Version
	caps    capability.List
	refs    *packp.AdvRefs
}

func NewStreamSession(conn Conn, service string) (*StreamSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StreamSession) Capabilities() *capability.List { _ = "STUB: not implemented"; return nil }

func (s *StreamSession) GetRemoteRefs(ctx context.Context, opts *GetRemoteRefsOptions) (*RemoteRefs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StreamSession) Fetch(ctx context.Context, st storage.Storer, req *FetchRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamSession) Push(ctx context.Context, st storage.Storer, req *PushRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamSession) Command(ctx context.Context, cmd string, req packp.CommandArgs, resp packp.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamSession) commandCapabilities() capability.List {
	_ = "STUB: not implemented"
	return *new(capability.List)
}

func (s *StreamSession) wrapStderr(err error) error { _ = "STUB: not implemented"; return nil }

func (s *StreamSession) Close() error { _ = "STUB: not implemented"; return nil }

func (s *StreamSession) Archive(ctx context.Context, req *ArchiveRequest) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

var (
	_ Session   = (*StreamSession)(nil)
	_ Archiver  = (*StreamSession)(nil)
	_ Commander = (*StreamSession)(nil)
)
