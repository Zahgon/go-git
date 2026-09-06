package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing"
)

func main() {
	CheckArgs("<path>", "<revision>")

	path := os.Args[1]
	revision := os.Args[2]

	r, err := git.PlainOpen(path)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git rev-parse %s", revision)

	h, err := r.ResolveRevision(plumbing.Revision(revision))

	CheckIfError(err)

	fmt.Println(h.String())
}
