package commitgraph

import (
	"io"

	"github.com/go-git/go-billy/v6"
)

func OpenChainFile(r io.Reader) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func OpenChainOrFileIndex(fs billy.Filesystem) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func OpenChainIndex(fs billy.Filesystem) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}
