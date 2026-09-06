package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v6/osfs"

	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/storage/filesystem"
	xworktree "github.com/go-git/go-git/v6/x/plumbing/worktree"
)

func main() {
	CheckArgs("<dotgit> <worktree>")
	path := os.Args[1]
	wtPath := os.Args[2]

	dotgitFs := osfs.New(filepath.Join(path, ".git"), osfs.WithBoundOS())
	dotgit := filesystem.NewStorageWithOptions(dotgitFs, nil, filesystem.Options{})

	w, err := xworktree.New(dotgit)
	CheckIfError(err)

	worktreeFs := osfs.New(wtPath)
	name := filepath.Base(wtPath)

	Info("git worktree add %s", wtPath)

	err = w.Add(worktreeFs, name)
	CheckIfError(err)

	Info("opening linked worktree at %q", wtPath)
	r, err := w.Open(worktreeFs)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	ref, err := r.Head()
	CheckIfError(err)

	c, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(c)
}
