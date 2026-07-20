package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)
	defer func() { _ = r.Close() }()

	ref, err := r.Head()
	CheckIfError(err)

	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
