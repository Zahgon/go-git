package dotgit

import (
	"bytes"
	"errors"
	"io"
	"os"
	"sync"
	"time"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/internal/packhandle"
	"github.com/go-git/go-git/v6/plumbing"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
	"github.com/go-git/go-git/v6/plumbing/format/packfile"
	"github.com/go-git/go-git/v6/x/fdpool"
)

const (
	packedRefsPath     = "packed-refs"
	configPath         = "config"
	configWorktreePath = "config.worktree"
	indexPath          = "index"
	shallowPath        = "shallow"
	modulePath         = "modules"
	objectsPath        = "objects"
	packPath           = "pack"
	refsPath           = "refs"
	branchesPath       = "branches"
	hooksPath          = "hooks"
	infoPath           = "info"
	remotesPath        = "remotes"
	logsPath           = "logs"
	worktreesPath      = "worktrees"
	alternatesPath     = "alternates"

	tmpPackedRefsPrefix = "._packed-refs"

	packPrefix = "pack-"
	packExt    = ".pack"

	promisorExt = ".promisor"
)

var (
	ErrNotFound = errors.New("path not found")

	ErrIdxNotFound = errors.New("idx file not found")

	ErrPackfileNotFound = errors.New("packfile not found")

	ErrConfigNotFound = errors.New("config file not found")

	ErrPackedRefsDuplicatedRef = errors.New("duplicated ref found in packed-ref file")

	ErrPackedRefsBadFormat = errors.New("malformed packed-ref")

	ErrSymRefTargetNotFound = errors.New("symbolic reference target not found")

	ErrIsDir = errors.New("reference path is a directory")

	ErrEmptyRefFile = errors.New("ref file is empty")

	ErrModuleNameEscape = errors.New("submodule name escapes modules/ directory")

	ErrReferenceNameEscape = errors.New("reference name escapes the reference storage")
)

func isPathSep(r rune) bool { _ = "STUB: not implemented"; return false }

func validReferenceName(name plumbing.ReferenceName) error { _ = "STUB: not implemented"; return nil }

type Options struct {
	ExclusiveAccess bool

	AlternatesFS billy.Filesystem

	ObjectFormat formatcfg.ObjectFormat

	ReadReverseIndex bool

	WriteReverseIndex bool

	Pool *fdpool.Pool
}

type DotGit struct {
	options Options
	fs      billy.Filesystem

	incomingOnce    sync.Once
	incomingDirName string

	objectList []plumbing.Hash
	objectMap  map[plumbing.Hash]struct{}
	packList   []plumbing.Hash
	packMap    map[plumbing.Hash]struct{}

	packHandlesMu sync.Mutex
	packHandles   map[plumbing.Hash]*packhandle.PackHandle
}

func New(fs billy.Filesystem) *DotGit { _ = "STUB: not implemented"; return nil }

func NewWithOptions(fs billy.Filesystem, o Options) *DotGit { _ = "STUB: not implemented"; return nil }

func (d *DotGit) Initialize() error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) Close() error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) ConfigWriter() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) Config() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ConfigWorktree() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ConfigWorktreeWriter() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) IndexWriter() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) Index() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) StatIndex() (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func (d *DotGit) ShallowWriter() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) Shallow() (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ReflogReader(name plumbing.ReferenceName) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ReflogWriter(name plumbing.ReferenceName) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) DeleteReflog(name plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) NewObjectPack() (*PackWriter, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *DotGit) NewPromisorObjectPack(marker string) (*PackWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) PromisorObjectPacks() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) hasPromisor(hash plumbing.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *DotGit) ObjectPacks() ([]plumbing.Hash, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *DotGit) objectPacks() ([]plumbing.Hash, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *DotGit) objectPackPath(hash plumbing.Hash, extension string) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DotGit) objectPackOpen(hash plumbing.Hash, extension string) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ObjectPack(hash plumbing.Hash) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ObjectPackIdx(hash plumbing.Hash) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ObjectPackRev(hash plumbing.Hash) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) OpenPackRev(hash plumbing.Hash) (idxfile.ReadAtCloser, error) {
	_ = "STUB: not implemented"
	return *new(idxfile.ReadAtCloser), nil
}

func (d *DotGit) generateInMemoryRev(h plumbing.Hash) (idxfile.ReadAtCloser, error) {
	_ = "STUB: not implemented"
	return *new(idxfile.ReadAtCloser), nil
}

type bytesReadAtCloser struct {
	*bytes.Reader
}

func newBytesReadAtCloser(data []byte) *bytesReadAtCloser { _ = "STUB: not implemented"; return nil }

func (b *bytesReadAtCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) packHandle(hash plumbing.Hash) (*packhandle.PackHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) walkPackHandles(fn func(*packhandle.PackHandle) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) CloseIdleDescriptors() error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) PackHandle(hash plumbing.Hash) (packfile.PackHandle, error) {
	_ = "STUB: not implemented"
	return *new(packfile.PackHandle), nil
}

type packHandleAdapter struct{ ph *packhandle.PackHandle }

func (a packHandleAdapter) OpenPackReader() (io.ReadSeekCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeekCloser), nil
}

func (a packHandleAdapter) OpenRandomReader() (packfile.RandomReader, error) {
	_ = "STUB: not implemented"
	return *new(packfile.RandomReader), nil
}

func (a packHandleAdapter) PackHash() (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (d *DotGit) DeleteOldObjectPackAndIndex(hash plumbing.Hash, t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) NewObject() (*ObjectWriter, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *DotGit) ObjectsWithPrefix(prefix []byte) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) Objects() ([]plumbing.Hash, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *DotGit) ForEachObjectHash(fun func(plumbing.Hash) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) forEachObjectHash(fun func(plumbing.Hash) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) cleanObjectList() { _ = "STUB: not implemented"; return }

func (d *DotGit) genObjectList() error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) hasObject(h plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) cleanPackList() error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) genPackList() error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) hasPack(h plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) objectPath(h plumbing.Hash) string { _ = "STUB: not implemented"; return "" }

func (d *DotGit) incomingObjectPath(h plumbing.Hash) string { _ = "STUB: not implemented"; return "" }

func (d *DotGit) hasIncomingObjects() bool { _ = "STUB: not implemented"; return false }

func (d *DotGit) Object(h plumbing.Hash) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) ObjectStat(h plumbing.Hash) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func (d *DotGit) ObjectDelete(h plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) readReferenceFrom(rd io.Reader, name string) (ref *plumbing.Reference, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) checkReferenceAndTruncate(f billy.File, old *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) SetRef(r, old *plumbing.Reference) error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) Refs() ([]*plumbing.Reference, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *DotGit) Ref(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) findPackedRefsInFile(f billy.File, recv refsRecv) error {
	_ = "STUB: not implemented"
	return nil
}

type refsRecv func(*plumbing.Reference) bool

func (d *DotGit) findPackedRefs(recv refsRecv) error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) packedRef(name plumbing.ReferenceName) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) RemoveRef(name plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

func refsRecvFunc(refs *[]*plumbing.Reference, seen map[plumbing.ReferenceName]bool) refsRecv {
	_ = "STUB: not implemented"
	return *new(refsRecv)
}

func (d *DotGit) addRefsFromPackedRefs(refs *[]*plumbing.Reference, seen map[plumbing.ReferenceName]bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) addRefsFromPackedRefsFile(refs *[]*plumbing.Reference, f billy.File, seen map[plumbing.ReferenceName]bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) openAndLockPackedRefs(doCreate bool) (
	pr billy.File, err error,
) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (d *DotGit) rewritePackedRefsWithoutRef(name plumbing.ReferenceName) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) processLine(line string) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) addRefsFromRefDir(refs *[]*plumbing.Reference, seen map[plumbing.ReferenceName]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) walkReferencesTree(refs *[]*plumbing.Reference, relPath []string, seen map[plumbing.ReferenceName]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) addRefFromHEAD(refs *[]*plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) readReferenceFile(path, name string) (ref *plumbing.Reference, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DotGit) CountLooseRefs() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *DotGit) PackRefs() (err error) { _ = "STUB: not implemented"; return nil }

func (d *DotGit) Module(name string) (billy.Filesystem, error) {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem), nil
}

func (d *DotGit) AddAlternate(remote string) error { _ = "STUB: not implemented"; return nil }

func (d *DotGit) Alternates() ([]*DotGit, error) { _ = "STUB: not implemented"; return nil, nil }

func alternatePathForFS(path, root string) string { _ = "STUB: not implemented"; return "" }

func (d *DotGit) Fs() billy.Filesystem { _ = "STUB: not implemented"; return *new(billy.Filesystem) }

func (d *DotGit) SetObjectFormat(of formatcfg.ObjectFormat) error {
	_ = "STUB: not implemented"
	return nil
}

func isHex(s string) bool { _ = "STUB: not implemented"; return false }

func isNum(b byte) bool { _ = "STUB: not implemented"; return false }

func isHexAlpha(b byte) bool { _ = "STUB: not implemented"; return false }

func incBytes(in []byte) (out []byte, overflow bool) { _ = "STUB: not implemented"; return nil, false }
