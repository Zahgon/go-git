package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing"
)

func main() {
	CheckArgs("<url>", "<directory>", "<commit>")
	url, directory, commit := os.Args[1], os.Args[2], os.Args[3]

	Info("git clone %s %s", url, directory)
	r, err := git.PlainClone(directory, &git.CloneOptions{
		URL: url,
	})

	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git show-ref --head HEAD")
	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())

	w, err := r.Worktree()
	CheckIfError(err)

	Info("git checkout %s", commit)
	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commit),
	})
	CheckIfError(err)

	Info("git show-ref --head HEAD")
	ref, err = r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())
}
