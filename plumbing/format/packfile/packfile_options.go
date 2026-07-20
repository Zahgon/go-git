package packfile

import (
	billy "github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
)

type PackfileOption func(*Packfile) //nolint:revive // stutters but is a well-established name

func WithCache(cache cache.Object) PackfileOption {
	_ = "STUB: not implemented"
	return *new(PackfileOption)
}

func WithIdx(idx idxfile.Index) PackfileOption {
	_ = "STUB: not implemented"
	return *new(PackfileOption)
}

func WithFs(fs billy.Filesystem) PackfileOption {
	_ = "STUB: not implemented"
	return *new(PackfileOption)
}

func WithObjectIDSize(sz int) PackfileOption {
	_ = "STUB: not implemented"
	return *new(PackfileOption)
}
