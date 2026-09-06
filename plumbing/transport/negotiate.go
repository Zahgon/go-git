package transport

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/storage"
)

const (
	initialFlush  = 16
	pipeSafeFlush = 32
	largeFlush    = 16384
	maxInVein     = 256
)

func nextFlush(statelessRPC bool, count int) int { _ = "STUB: not implemented"; return 0 }

func applyServerACKs(
	statelessRPC bool,
	acks []packp.ACK,
	common map[plumbing.Hash]struct{},
	statelessCommon *[]plumbing.Hash,
	gotContinue *bool,
	gotReady *bool,
	inVein *int,
) {
	_ = "STUB: not implemented"
	return
}

func NegotiatePack(
	ctx context.Context,
	st storage.Storer,
	caps capability.List,
	statelessRPC bool,
	reader io.Reader,
	writer io.WriteCloser,
	req *FetchRequest,
) (shallowInfo *packp.ShallowUpdate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isSubset(needle, haystack []plumbing.Hash) bool { _ = "STUB: not implemented"; return false }

func readShallows(
	statelessRPC bool,
	r io.Reader,
	req *FetchRequest,
	shallowInfo **packp.ShallowUpdate,
	firstRound bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ReconcileObjectFormatV2(st storage.Storer, caps capability.List) error {
	_ = "STUB: not implemented"
	return nil
}
