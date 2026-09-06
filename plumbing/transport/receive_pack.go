package transport

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/storage"
)

type ReceivePackRequest struct {
	GitProtocol   string
	AdvertiseRefs bool
	StatelessRPC  bool

	Hooks ReceivePackHooks
}

type ReceivePackHooks struct {
	PreReceive func(context.Context, *PreReceiveInfo) error

	PostReceive func(context.Context, *PostReceiveInfo) error
}

type PreReceiveInfo struct {
	Storer storage.Storer

	Commands []*packp.Command

	PushOptions []string

	Progress io.Writer
}

type PostReceiveInfo struct {
	Storer storage.Storer

	Commands []*packp.Command

	PushOptions []string

	Progress io.Writer
}

func ReceivePack(
	ctx context.Context,
	st storage.Storer,
	r io.ReadCloser,
	w io.WriteCloser,
	opts *ReceivePackRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

type sidebandProgress struct{ mux *sideband.Muxer }

func (p sidebandProgress) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func closeWriter(w io.WriteCloser) error { _ = "STUB: not implemented"; return nil }

func sendReportStatus(w io.WriteCloser, unpackErr error, cmdStatus map[plumbing.ReferenceName]error) error {
	_ = "STUB: not implemented"
	return nil
}

func setStatus(cmdStatus map[plumbing.ReferenceName]error, firstErr *error, ref plumbing.ReferenceName, err error) {
	_ = "STUB: not implemented"
	return
}

func referenceExists(s storer.ReferenceStorer, n plumbing.ReferenceName) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func updateReferences(st storage.Storer, req *packp.UpdateRequests, cmdStatus map[plumbing.ReferenceName]error, firstErr *error) {
	_ = "STUB: not implemented"
	return
}
