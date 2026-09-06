package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing"
)

func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	ref, err := r.Head()
	CheckIfError(err)

	tags, err := r.Tags()
	CheckIfError(err)

	err = tags.ForEach(func(t *plumbing.Reference) error {

		revHash, err := r.ResolveRevision(plumbing.Revision(t.Name()))
		CheckIfError(err)
		if *revHash == ref.Hash() {
			fmt.Printf("Found tag %s with hash %s pointing to HEAD %s\n", t.Name().Short(), revHash, ref.Hash())
		}
		return nil
	})
	CheckIfError(err)
}
