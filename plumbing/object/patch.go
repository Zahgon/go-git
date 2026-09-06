package object

import (
	"context"
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/filemode"
	fdiff "github.com/go-git/go-git/v6/plumbing/format/diff"
)

var ErrCanceled = errors.New("operation canceled")

func getPatch(message string, changes ...*Change) (*Patch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPatchContext(ctx context.Context, message string, changes ...*Change) (*Patch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filePatchWithContext(ctx context.Context, c *Change) (fdiff.FilePatch, error) {
	_ = "STUB: not implemented"
	return *new(fdiff.FilePatch), nil
}

func isSubmodule(e ChangeEntry) bool { _ = "STUB: not implemented"; return false }

func submoduleContent(e ChangeEntry) string { _ = "STUB: not implemented"; return "" }

func submoduleFilePatch(ctx context.Context, c *Change) (fdiff.FilePatch, error) {
	_ = "STUB: not implemented"
	return *new(fdiff.FilePatch), nil
}

func fileContent(f *File) (content string, isBinary bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

type Patch struct {
	message     string
	filePatches []fdiff.FilePatch
}

func (p *Patch) FilePatches() []fdiff.FilePatch { _ = "STUB: not implemented"; return nil }

func (p *Patch) Message() string { _ = "STUB: not implemented"; return "" }

func (p *Patch) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (p *Patch) Stats() FileStats { _ = "STUB: not implemented"; return *new(FileStats) }

func (p *Patch) String() string { _ = "STUB: not implemented"; return "" }

type changeEntryWrapper struct {
	ce ChangeEntry
}

func (f *changeEntryWrapper) Hash() plumbing.Hash {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash)
}

func (f *changeEntryWrapper) Mode() filemode.FileMode {
	_ = "STUB: not implemented"
	return *new(filemode.FileMode)
}

func (f *changeEntryWrapper) Path() string { _ = "STUB: not implemented"; return "" }

func (f *changeEntryWrapper) Empty() bool { _ = "STUB: not implemented"; return false }

type textFilePatch struct {
	chunks   []fdiff.Chunk
	from, to ChangeEntry
}

func (tf *textFilePatch) Files() (from, to fdiff.File) {
	_ = "STUB: not implemented"
	return *new(fdiff.File), *new(fdiff.File)
}

func (tf *textFilePatch) IsBinary() bool { _ = "STUB: not implemented"; return false }

func (tf *textFilePatch) Chunks() []fdiff.Chunk { _ = "STUB: not implemented"; return nil }

type textChunk struct {
	content string
	op      fdiff.Operation
}

func (t *textChunk) Content() string { _ = "STUB: not implemented"; return "" }

func (t *textChunk) Type() fdiff.Operation { _ = "STUB: not implemented"; return *new(fdiff.Operation) }

type FileStat struct {
	Name     string
	Addition int
	Deletion int
}

func (fs FileStat) String() string { _ = "STUB: not implemented"; return "" }

type FileStats []FileStat

func (fileStats FileStats) String() string { _ = "STUB: not implemented"; return "" }

func printStat(fileStats []FileStat) string { _ = "STUB: not implemented"; return "" }

func getFileStatsFromFilePatches(filePatches []fdiff.FilePatch) FileStats {
	_ = "STUB: not implemented"
	return *new(FileStats)
}
