package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func main() {
	CheckArgs("<path>", "<file_to_blame>")
	url := os.Args[1]
	path := os.Args[2]

	Info("git open %s", url)
	r, err := git.PlainOpen(url)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	ref, err := r.Head()
	CheckIfError(err)

	c, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	Info("git blame %s", path)

	br, err := git.Blame(c, path)
	CheckIfError(err)

	fmt.Printf("%s", br.String())
}
