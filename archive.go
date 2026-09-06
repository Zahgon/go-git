package git

import (
	"context"
	"io"
)

func ArchiveRemote(url string, o *ArchiveOptions) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func ArchiveRemoteContext(ctx context.Context, url string, o *ArchiveOptions) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
