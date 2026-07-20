package git

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

type signableObject interface {
	EncodeWithoutSignature(o plumbing.EncodedObject) error
}

type Signer interface {
	Sign(ctx context.Context, message io.Reader) ([]byte, error)
}

func signObject(signer Signer, obj signableObject) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
