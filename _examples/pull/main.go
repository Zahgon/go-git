package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	w, err := r.Worktree()
	CheckIfError(err)

	Info("git pull origin")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	CheckIfError(err)

	ref, err := r.Head()
	CheckIfError(err)
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
