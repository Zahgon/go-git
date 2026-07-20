package transport

import (
	"context"
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/protocol"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/storage"
)

var ErrUpdateReference = errors.New("failed to update ref")

func AdvertiseRefs(
	_ context.Context,
	st storage.Storer,
	w io.Writer,
	service string,
	smart bool,
	version protocol.Version,
) error {
	_ = "STUB: not implemented"
	return nil
}

func AdvertiseCapabilities(_ context.Context, st storage.Storer, w io.Writer, service string) error {
	_ = "STUB: not implemented"
	return nil
}

func serverV2Capabilities(st storage.Storer) capability.List {
	_ = "STUB: not implemented"
	return *new(capability.List)
}

func objectFormat(st storage.Storer) config.ObjectFormat {
	_ = "STUB: not implemented"
	return *new(config.ObjectFormat)
}

func addReferences(st storage.Storer, ar *packp.AdvRefs, addHead bool) error {
	_ = "STUB: not implemented"
	return nil
}
