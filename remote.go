package git

import (
	"context"
	"errors"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/client"
	"github.com/go-git/go-git/v6/plumbing/protocol"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/go-git/go-git/v6/storage"
	"github.com/go-git/go-git/v6/storage/memory"
)

var (
	NoErrAlreadyUpToDate     = errors.New("already up-to-date") //nolint:staticcheck,revive // sentinel value, not an error
	ErrDeleteRefNotSupported = errors.New("server does not support delete-refs")
	ErrForceNeeded           = errors.New("some refs were not updated")
	ErrExactSHA1NotSupported = errors.New("server does not support exact SHA1 refspec")
	ErrEmptyUrls             = errors.New("URLs cannot be empty")
	ErrRemoteRefNotFound     = errors.New("couldn't find remote ref")
)

const (
	maxHavesToVisitPerRef = 100

	peeledSuffix = "^{}"
)

type Remote struct {
	c *config.RemoteConfig
	s storage.Storer
}

func NewRemote(s storage.Storer, c *config.RemoteConfig) *Remote {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) Config() *config.RemoteConfig { _ = "STUB: not implemented"; return nil }

func (r *Remote) String() string { _ = "STUB: not implemented"; return "" }

func (r *Remote) Push(o *PushOptions) error { _ = "STUB: not implemented"; return nil }

func (r *Remote) PushContext(ctx context.Context, o *PushOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) sendPack(ctx context.Context, sess transport.Session, remoteRefs storer.ReferenceStorer, o *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) useRefDeltas(ar *packp.AdvRefs) bool { _ = "STUB: not implemented"; return false }

func (r *Remote) addReachableTags(localRefs []*plumbing.Reference, remoteRefs storer.ReferenceStorer, cmds *[]*packp.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) updateRemoteReferenceStorage(
	cmds []*packp.Command,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) FetchContext(ctx context.Context, o *FetchOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchRefPrefixes(specs []config.RefSpec, tags plumbing.TagMode) []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) Fetch(o *FetchOptions) error { _ = "STUB: not implemented"; return nil }

func (r *Remote) fetch(ctx context.Context, o *FetchOptions) (sto storer.ReferenceStorer, err error) {
	_ = "STUB: not implemented"
	return *new(storer.ReferenceStorer), nil
}

func referenceStorageFromRefs(refs []*plumbing.Reference, filterPeeled bool) memory.ReferenceStorage {
	_ = "STUB: not implemented"
	return *new(memory.ReferenceStorage)
}

func depthChanged(before []plumbing.Hash, s storage.Storer) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func newClient(rawURL string, opts []client.Option) (*client.Client, *transport.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *Remote) transportProtocol() protocol.Version {
	_ = "STUB: not implemented"
	return *new(protocol.Version)
}

func (r *Remote) pruneRemotes(specs []config.RefSpec, localRefs []*plumbing.Reference, remoteRefs storer.ReferenceStorer) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Remote) addReferencesToUpdate(
	refspecs []config.RefSpec,
	localRefs []*plumbing.Reference,
	remoteRefs storer.ReferenceStorer,
	cmds *[]*packp.Command,
	prune bool,
	forceWithLease *ForceWithLease,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) addOrUpdateReferences(
	rs config.RefSpec,
	localRefs []*plumbing.Reference,
	refsDict map[string]*plumbing.Reference,
	remoteRefs storer.ReferenceStorer,
	cmds *[]*packp.Command,
	forceWithLease *ForceWithLease,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) deleteReferences(rs config.RefSpec,
	remoteRefs storer.ReferenceStorer,
	refsDict map[string]*plumbing.Reference,
	cmds *[]*packp.Command,
	prune bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) addObject(rs config.RefSpec,
	remoteRefs storer.ReferenceStorer, localObject plumbing.Hash,
	cmds *[]*packp.Command,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) addReferenceIfRefSpecMatches(rs config.RefSpec,
	remoteRefs storer.ReferenceStorer, localRef *plumbing.Reference,
	cmds *[]*packp.Command, forceWithLease *ForceWithLease,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) checkForceWithLease(localRef *plumbing.Reference, cmd *packp.Command, forceWithLease *ForceWithLease) error {
	_ = "STUB: not implemented"
	return nil
}

func getRemoteRefsFromStorer(remoteRefStorer storer.ReferenceStorer) (
	map[plumbing.Hash]bool, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getHavesFromRef(
	ref *plumbing.Reference,
	remoteRefs map[plumbing.Hash]bool,
	s storage.Storer,
	haves map[plumbing.Hash]bool,
	depth int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func getHaves(
	localRefs []*plumbing.Reference,
	remoteRefStorer storer.ReferenceStorer,
	s storage.Storer,
	depth int,
) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const refspecAllTags = "+refs/tags/*:refs/tags/*"

func calculateRefs(
	spec []config.RefSpec,
	remoteRefs storer.ReferenceStorer,
	tagMode plumbing.TagMode,
) (memory.ReferenceStorage, [][]*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return *new(memory.ReferenceStorage), nil, nil
}

func doCalculateRefs(
	s config.RefSpec,
	remoteRefs storer.ReferenceStorer,
	refs memory.ReferenceStorage,
) ([]*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getWants(localStorer storage.Storer, refs memory.ReferenceStorage, depth int) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func objectExists(s storer.EncodedObjectStorer, h plumbing.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkFastForwardUpdate(s storer.EncodedObjectStorer, remoteRefs storer.ReferenceStorer, cmd *packp.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func isFastForward(s storer.EncodedObjectStorer, old, newHash plumbing.Hash, shallows []plumbing.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Remote) isSupportedRefSpec(refs []config.RefSpec, caps *capability.List) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) updateLocalReferenceStorage(
	specs []config.RefSpec,
	fetchedRefs, remoteRefs memory.ReferenceStorage,
	specToRefs [][]*plumbing.Reference,
	tagMode plumbing.TagMode,
	force bool,
) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Remote) buildFetchedTags(refs memory.ReferenceStorage) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Remote) ListContext(ctx context.Context, o *ListOptions) (rfs []*plumbing.Reference, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Remote) List(o *ListOptions) (rfs []*plumbing.Reference, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Remote) list(ctx context.Context, o *ListOptions) (rfs []*plumbing.Reference, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func objectsToPush(commands []*packp.Command) []plumbing.Hash {
	_ = "STUB: not implemented"
	return nil
}

func referencesToHashes(refs storer.ReferenceStorer) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pushHashes(
	ctx context.Context,
	sess transport.Session,
	s storage.Storer,
	cmds []*packp.Command,
	hs []plumbing.Hash,
	allDelete bool,
	o *PushOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) checkRequireRemoteRefs(requires []config.RefSpec, remoteRefs storer.ReferenceStorer) error {
	_ = "STUB: not implemented"
	return nil
}
