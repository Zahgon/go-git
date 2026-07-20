package git

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/client"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/storage"
)

const GitDirName = ".git"

var (
	ErrBranchExists = errors.New("branch already exists")

	ErrBranchNotFound = errors.New("branch not found")

	ErrTagExists = errors.New("tag already exists")

	ErrTagNotFound = errors.New("tag not found")

	ErrFetching = errors.New("unable to fetch packfile")

	ErrInvalidReference = errors.New("invalid reference, should be a tag or a branch")

	ErrRepositoryNotExists = errors.New("repository does not exist")

	ErrRepositoryIncomplete = errors.New("repository's commondir path does not exist")

	ErrRemoteNotFound = errors.New("remote not found")

	ErrRemoteExists = errors.New("remote already exists")

	ErrAnonymousRemoteName = errors.New("anonymous remote name must be 'anonymous'")

	ErrWorktreeNotProvided = errors.New("worktree should be provided")

	ErrIsBareRepository = errors.New("worktree not available in a bare repository")

	ErrUnableToResolveCommit = errors.New("unable to resolve commit")

	ErrPackedObjectsNotSupported = errors.New("packed objects not supported")

	ErrAlternatePathNotSupported = errors.New("alternate path must use the file scheme")

	ErrUnsupportedMergeStrategy = errors.New("unsupported merge strategy")

	ErrFastForwardMergeNotPossible = errors.New("not possible to fast-forward merge changes")

	ErrTargetDirNotEmpty = errors.New("destination path already exists and is not empty")
)

type Repository struct {
	Storer storage.Storer

	r  map[string]*Remote
	wt billy.Filesystem
}

type initOptions struct {
	defaultBranch plumbing.ReferenceName
	workTree      billy.Filesystem
	objectFormat  formatcfg.ObjectFormat
	partialInit   bool
}

func newInitOptions() initOptions { _ = "STUB: not implemented"; return *new(initOptions) }

type InitOption func(*initOptions)

func WithDefaultBranch(b plumbing.ReferenceName) InitOption {
	_ = "STUB: not implemented"
	return *new(InitOption)
}

func WithWorkTree(worktree billy.Filesystem) InitOption {
	_ = "STUB: not implemented"
	return *new(InitOption)
}

func WithObjectFormat(of formatcfg.ObjectFormat) InitOption {
	_ = "STUB: not implemented"
	return *new(InitOption)
}

func withPartialInit() InitOption { _ = "STUB: not implemented"; return *new(InitOption) }

func Init(s storage.Storer, opts ...InitOption) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initStorer(s storer.Storer) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) setWorktreeAndStoragePaths() error { _ = "STUB: not implemented"; return nil }

func createDotGitFile(worktree, storage billy.Filesystem) error {
	_ = "STUB: not implemented"
	return nil
}

func setConfigWorktree(r *Repository, worktree, storage billy.Filesystem) error {
	_ = "STUB: not implemented"
	return nil
}

func Open(s storage.Storer, worktree billy.Filesystem) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Clone(s storage.Storer, worktree billy.Filesystem, o *CloneOptions) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CloneContext(
	ctx context.Context, s storage.Storer, worktree billy.Filesystem, o *CloneOptions,
) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PlainInit(path string, isBare bool, options ...InitOption) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) setInvalidHEAD() (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PlainOpen(path string) (*Repository, error) { _ = "STUB: not implemented"; return nil, nil }

func PlainOpenWithOptions(path string, o *PlainOpenOptions) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dotGitToOSFilesystems(path string, detect bool) (dot, wt billy.Filesystem, err error) {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem), *new(billy.Filesystem), nil
}

func dotGitFileToOSFilesystem(path string, fs billy.Filesystem) (bfs billy.Filesystem, err error) {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem), nil
}

func dotGitCommonDirectory(fs billy.Filesystem) (commonDir billy.Filesystem, err error) {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem), nil
}

func PlainClone(path string, o *CloneOptions) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PlainCloneContext(ctx context.Context, path string, o *CloneOptions) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newRepository(s storage.Storer, worktree billy.Filesystem) *Repository {
	_ = "STUB: not implemented"
	return nil
}

func checkTargetDirIsEmpty(path string) (empty bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Repository) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Repository) Config() (*config.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Repository) SetConfig(cfg *config.Config) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) ConfigScoped(scope config.Scope) (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Remote(name string) (*Remote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Remotes() ([]*Remote, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Repository) CreateRemote(c *config.RemoteConfig) (*Remote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) CreateRemoteAnonymous(c *config.RemoteConfig) (*Remote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) DeleteRemote(name string) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) Branch(name string) (*config.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) CreateBranch(c *config.Branch) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) DeleteBranch(name string) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) CreateTag(name string, hash plumbing.Hash, opts *CreateTagOptions) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) createTagObject(name string, hash plumbing.Hash, opts *CreateTagOptions) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (r *Repository) buildTagSignature(tag *object.Tag, signer Signer) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *Repository) Tag(name string) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) DeleteTag(name string) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) resolveToCommitHash(h plumbing.Hash) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (r *Repository) clone(ctx context.Context, o *CloneOptions) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	refspecTag              = "+refs/tags/%s:refs/tags/%[1]s"
	refspecSingleBranch     = "+refs/heads/%s:refs/remotes/%s/%[1]s"
	refspecSingleBranchHEAD = "+HEAD:refs/remotes/%s/HEAD"
)

func (r *Repository) cloneRefSpec(o *CloneOptions) []config.RefSpec {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) setIsBare(isBare bool) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) updateRemoteConfigIfNeeded(o *CloneOptions, c *config.RemoteConfig, _ *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) fetchAndUpdateReferences(
	ctx context.Context, o *FetchOptions, ref plumbing.ReferenceName,
) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) updateReferences(spec []config.RefSpec,
	resolvedRef *plumbing.Reference,
) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Repository) calculateRemoteHeadReference(spec []config.RefSpec,
	resolvedHead *plumbing.Reference,
) []*plumbing.Reference {
	_ = "STUB: not implemented"
	return nil
}

func checkAndUpdateReferenceStorerIfNeeded(
	s storer.ReferenceStorer, r, old *plumbing.Reference) (
	updated bool, err error,
) {
	_ = "STUB: not implemented"
	return false, nil
}

func updateReferenceStorerIfNeeded(
	s storer.ReferenceStorer, r *plumbing.Reference,
) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Repository) Fetch(o *FetchOptions) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) FetchContext(ctx context.Context, o *FetchOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) Push(o *PushOptions) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) PushContext(ctx context.Context, o *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type ArchiveOptions struct {
	Format string

	Prefix string

	Treeish string

	Paths []string

	ClientOptions []client.Option

	Progress sideband.Progress
}

func (o *ArchiveOptions) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *Repository) Archive(o *ArchiveOptions) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (r *Repository) ArchiveContext(ctx context.Context, o *ArchiveOptions) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (r *Repository) Log(o *LogOptions) (object.CommitIter, error) {
	_ = "STUB: not implemented"
	return *new(object.CommitIter), nil
}

func (r *Repository) log(from plumbing.Hash, commitIterFunc func(*object.Commit) object.CommitIter) (object.CommitIter, error) {
	_ = "STUB: not implemented"
	return *new(object.CommitIter), nil
}

func (r *Repository) logAll(commitIterFunc func(*object.Commit) object.CommitIter) (object.CommitIter, error) {
	_ = "STUB: not implemented"
	return *new(object.CommitIter), nil
}

func (*Repository) logWithFile(fileName string, commitIter object.CommitIter, checkParent bool) object.CommitIter {
	_ = "STUB: not implemented"
	return *new(object.CommitIter)
}

func (*Repository) logWithPathFilter(pathFilter func(string) bool, commitIter object.CommitIter, checkParent bool) object.CommitIter {
	_ = "STUB: not implemented"
	return *new(object.CommitIter)
}

func (*Repository) logWithLimit(commitIter object.CommitIter, limitOptions object.LogLimitOptions) object.CommitIter {
	_ = "STUB: not implemented"
	return *new(object.CommitIter)
}

func commitIterFunc(order LogOrder) func(c *object.Commit) object.CommitIter {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) Tags() (storer.ReferenceIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.ReferenceIter), nil
}

func (r *Repository) Branches() (storer.ReferenceIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.ReferenceIter), nil
}

func (r *Repository) Notes() (storer.ReferenceIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.ReferenceIter), nil
}

func (r *Repository) TreeObject(h plumbing.Hash) (*object.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) TreeObjects() (*object.TreeIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) CommitObject(h plumbing.Hash) (*object.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) CommitObjects() (object.CommitIter, error) {
	_ = "STUB: not implemented"
	return *new(object.CommitIter), nil
}

func (r *Repository) BlobObject(h plumbing.Hash) (*object.Blob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) BlobObjects() (*object.BlobIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) TagObject(h plumbing.Hash) (*object.Tag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) TagObjects() (*object.TagIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Object(t plumbing.ObjectType, h plumbing.Hash) (object.Object, error) {
	_ = "STUB: not implemented"
	return *new(object.Object), nil
}

func (r *Repository) Objects() (*object.ObjectIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Head() (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Reference(name plumbing.ReferenceName, resolved bool) (
	*plumbing.Reference, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) References() (storer.ReferenceIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.ReferenceIter), nil
}

func (r *Repository) Worktree() (*Worktree, error) { _ = "STUB: not implemented"; return nil, nil }

func expandRef(s storer.ReferenceStorer, ref plumbing.ReferenceName) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) ResolveRevision(in plumbing.Revision) (*plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) resolveHashPrefix(hashStr string) []plumbing.Hash {
	_ = "STUB: not implemented"
	return nil
}

type RepackConfig struct {
	UseRefDeltas bool

	OnlyDeletePacksOlderThan time.Time
}

func (r *Repository) RepackObjects(cfg *RepackConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) Merge(ref plumbing.Reference, opts MergeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) createNewObjectPack(cfg *RepackConfig) (h plumbing.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func expandPartialHash(st storer.EncodedObjectStorer, prefix []byte) (hashes []plumbing.Hash) {
	_ = "STUB: not implemented"
	return nil
}
