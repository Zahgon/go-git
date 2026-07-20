package merkletrie

import (
	"context"
	"errors"

	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

var ErrCanceled = errors.New("operation canceled")

func DiffTree(
	fromTree,
	toTree noder.Noder,
	hashEqual noder.Equal,
) (Changes, error) {
	_ = "STUB: not implemented"
	return *new(Changes), nil
}

func DiffTreeContext(ctx context.Context, fromTree, toTree noder.Noder,
	hashEqual noder.Equal,
) (Changes, error) {
	_ = "STUB: not implemented"
	return *new(Changes), nil
}

func diffNodes(changes *Changes, ii *doubleIter) error { _ = "STUB: not implemented"; return nil }

func diffNodesSameName(changes *Changes, ii *doubleIter) error {
	_ = "STUB: not implemented"
	return nil
}

func diffDirs(changes *Changes, ii *doubleIter) error { _ = "STUB: not implemented"; return nil }
