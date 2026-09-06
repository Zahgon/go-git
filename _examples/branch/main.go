package main

import (
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing"
)

func main() {
	CheckArgs("<url>", "<directory>")
	url, directory := os.Args[1], os.Args[2]

	Info("git clone %s %s", url, directory)
	r, err := git.PlainClone(directory, &git.CloneOptions{
		URL: url,
	})
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git branch my-branch")

	headRef, err := r.Head()
	CheckIfError(err)

	ref := plumbing.NewHashReference("refs/heads/my-branch", headRef.Hash())

	err = r.Storer.SetReference(ref)
	CheckIfError(err)

	Info("git branch -D my-branch")
	err = r.Storer.RemoveReference(ref.Name())
	CheckIfError(err)
}
