package main

import (
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func main() {
	CheckArgs("<url>", "<directory>", "<submodule>")
	url := os.Args[1]
	directory := os.Args[2]
	submodule := os.Args[3]

	Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)
	defer func() { _ = r.Close() }()

	w, err := r.Worktree()
	if err != nil {
		CheckIfError(err)
	}

	sub, err := w.Submodule(submodule)
	if err != nil {
		CheckIfError(err)
	}

	sr, err := sub.Repository()
	if err != nil {
		CheckIfError(err)
	}

	sw, err := sr.Worktree()
	if err != nil {
		CheckIfError(err)
	}

	Info("git submodule update --remote")
	err = sw.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil {
		CheckIfError(err)
	}
}
