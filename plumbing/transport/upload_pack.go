package transport

import (
	"bufio"
	"context"
	"io"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/storage"
)

type UploadPackRequest struct {
	GitProtocol   string
	AdvertiseRefs bool
	StatelessRPC  bool

	SkipDeltaCompression bool
}

func UploadPack(
	ctx context.Context,
	st storage.Storer,
	r io.ReadCloser,
	w io.WriteCloser,
	opts *UploadPackRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

func objectsToUpload(st storage.Storer, wants, haves []plumbing.Hash) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getShallowCommits(st storage.Storer, heads []plumbing.Hash, depth int, upd *packp.ShallowUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

func shallowFrontierDepth(st storage.Storer, heads, shallows []plumbing.Hash) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func serveUploadPackV2(ctx context.Context, st storage.Storer, rd *bufio.Reader, w io.WriteCloser, opts *UploadPackRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func serveLsRefsV2(_ context.Context, st storage.Storer, w io.Writer, args *packp.LsRefsArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func refMatchesAnyPrefix(name string, prefixes []string) bool {
	_ = "STUB: not implemented"
	return false
}

func writeV2Ref(w io.Writer, st storage.Storer, r *plumbing.Reference, symrefs, peel bool) error {
	_ = "STUB: not implemented"
	return nil
}

func peelToNonTag(st storage.Storer, h plumbing.Hash) (plumbing.Hash, bool) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), false
}

func serveFetchV2(_ context.Context, st storage.Storer, w io.WriteCloser, args *packp.FetchArgs, opts *UploadPackRequest) (concluded bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func hashDifference(a, b []plumbing.Hash) []plumbing.Hash { _ = "STUB: not implemented"; return nil }

func unshallowedCommits(clientShallows, newBoundary, newView []plumbing.Hash) []plumbing.Hash {
	_ = "STUB: not implemented"
	return nil
}

func resolveDeepenNot(st storage.Storer, refs []string) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func reachableCommits(st storage.Storer, tips []plumbing.Hash) (map[plumbing.Hash]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getShallowCommitsByRevList(st storage.Storer, heads []plumbing.Hash, since time.Time, notTips []plumbing.Hash, upd *packp.ShallowUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

func includeReachableTags(st storage.Storer, objs []plumbing.Hash) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type shallowBoundaryStorer struct {
	storage.Storer
	boundary []plumbing.Hash
}

func (s *shallowBoundaryStorer) Shallow() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wantsReachableFromHaves(st storage.Storer, wants, commonHaves []plumbing.Hash) bool {
	_ = "STUB: not implemented"
	return false
}

func peelToCommit(st storage.Storer, h plumbing.Hash) (*object.Commit, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
