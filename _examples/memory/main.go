package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-billy/v6/memfs"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/storage/memory"
)

func main() {
	CheckArgs("<url>")
	url := os.Args[1]

	Info("git clone %s", url)

	wt := memfs.New()
	storer := memory.NewStorage()
	r, err := git.Clone(storer, wt, &git.CloneOptions{
		URL: url,
	})

	CheckIfError(err)
	defer func() { _ = r.Close() }()

	ref, err := r.Head()
	CheckIfError(err)

	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
