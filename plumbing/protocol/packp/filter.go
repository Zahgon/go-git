package packp

import (
	"errors"

	"github.com/go-git/go-git/v6/plumbing"
)

var ErrUnsupportedObjectFilterType = errors.New("unsupported object filter type")

type Filter string

type BlobLimitPrefix string

const (
	BlobLimitPrefixNone BlobLimitPrefix = ""
	BlobLimitPrefixKibi BlobLimitPrefix = "k"
	BlobLimitPrefixMebi BlobLimitPrefix = "m"
	BlobLimitPrefixGibi BlobLimitPrefix = "g"
)

func FilterBlobNone() Filter { _ = "STUB: not implemented"; return *new(Filter) }

func FilterBlobLimit(n uint64, prefix BlobLimitPrefix) Filter {
	_ = "STUB: not implemented"
	return *new(Filter)
}

func FilterTreeDepth(depth uint64) Filter { _ = "STUB: not implemented"; return *new(Filter) }

func FilterObjectType(t plumbing.ObjectType) (Filter, error) {
	_ = "STUB: not implemented"
	return *new(Filter), nil
}

func FilterCombine(filters ...Filter) Filter { _ = "STUB: not implemented"; return *new(Filter) }
