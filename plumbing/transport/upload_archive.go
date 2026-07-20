package transport

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v6/storage"
)

type UploadArchiveRequest struct{}

func UploadArchive(
	ctx context.Context,
	st storage.Storer,
	r io.ReadCloser,
	w io.WriteCloser,
	_ *UploadArchiveRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

const maxArchiveArgs = 64

func readArchiveArgs(r io.Reader) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func writeNACK(w io.Writer, reason string) { _ = "STUB: not implemented"; return }

func muxError(mux *sideband.Muxer, w io.Writer, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func unsupportedArchiveFeature(arg string) string { _ = "STUB: not implemented"; return "" }
