package transport

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/storage"
)

func FetchPack(
	ctx context.Context,
	st storage.Storer,
	caps capability.List,
	packf io.ReadCloser,
	shallowInfo *packp.ShallowUpdate,
	req *FetchRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

func updateShallow(st storage.Storer, shallowInfo *packp.ShallowUpdate) error {
	_ = "STUB: not implemented"
	return nil
}
