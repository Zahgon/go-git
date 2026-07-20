package transport

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol/packp/sideband"
)

type ArchiveRequest struct {
	Args []string

	Progress sideband.Progress
}

type Archiver interface {
	Archive(ctx context.Context, req *ArchiveRequest) (io.ReadCloser, error)
}

func Archive(ctx context.Context, w io.WriteCloser, r io.ReadCloser, req *ArchiveRequest) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
