package object

import (
	"context"
)

func DiffTree(a, b *Tree) (Changes, error) { _ = "STUB: not implemented"; return *new(Changes), nil }

func DiffTreeContext(ctx context.Context, a, b *Tree) (Changes, error) {
	_ = "STUB: not implemented"
	return *new(Changes), nil
}

type DiffTreeOptions struct {
	DetectRenames bool

	RenameScore uint

	RenameLimit uint

	OnlyExactRenames bool
}

var DefaultDiffTreeOptions = &DiffTreeOptions{
	DetectRenames:    true,
	RenameScore:      60,
	RenameLimit:      0,
	OnlyExactRenames: false,
}

func DiffTreeWithOptions(
	ctx context.Context,
	a, b *Tree,
	opts *DiffTreeOptions,
) (Changes, error) {
	_ = "STUB: not implemented"
	return *new(Changes), nil
}
