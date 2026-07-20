package merkletrie

import (
	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

type doubleIter struct {
	from struct {
		iter    *Iter
		current noder.Path
	}
	to struct {
		iter    *Iter
		current noder.Path
	}
	hashEqual noder.Equal
}

func newDoubleIter(from, to noder.Noder, hashEqual noder.Equal) (
	*doubleIter, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func turnEOFIntoNil(e error) error { _ = "STUB: not implemented"; return nil }

func (d *doubleIter) nextBoth() error { _ = "STUB: not implemented"; return nil }

func (d *doubleIter) nextFrom() (err error) { _ = "STUB: not implemented"; return nil }

func (d *doubleIter) nextTo() (err error) { _ = "STUB: not implemented"; return nil }

func (d *doubleIter) stepBoth() (err error) { _ = "STUB: not implemented"; return nil }

func (d *doubleIter) remaining() remaining { _ = "STUB: not implemented"; return *new(remaining) }

type remaining int

const (
	noMoreNoders remaining = iota
	onlyToRemains
	onlyFromRemains
	bothHaveNodes
)

func (d *doubleIter) compare() (s comparison, err error) {
	_ = "STUB: not implemented"
	return *new(comparison), nil
}

type comparison struct {
	sameHash bool

	bothAreFiles bool

	fileAndDir bool

	bothAreDirs bool

	fromIsEmptyDir bool

	toIsEmptyDir bool
}
