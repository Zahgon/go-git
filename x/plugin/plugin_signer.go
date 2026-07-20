package plugin

import (
	"context"
	"io"
)

const objectSignerPlugin Name = "object-signer"

var objectSigner = newKey[Signer](objectSignerPlugin)

type Signer interface {
	Sign(ctx context.Context, message io.Reader) ([]byte, error)
}

func ObjectSigner() key[Signer] {
	_ = "STUB: not implemented" //nolint:revive // intentional unexported return type
	return nil
}
