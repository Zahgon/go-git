package transport

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/storage"
)

func SendPack(
	ctx context.Context,
	_ storage.Storer,
	caps capability.List,
	writer io.WriteCloser,
	reader io.ReadCloser,
	req *PushRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

func buildUpdateRequests(caps capability.List, req *PushRequest) *packp.UpdateRequests {
	_ = "STUB: not implemented"
	return nil
}
