package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing/object"
)

func main() {
	CheckArgs("<directory>")
	directory := os.Args[1]

	r, err := git.PlainOpen(directory)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	w, err := r.Worktree()
	CheckIfError(err)

	Info("echo \"hello world!\" > example-git-file")
	filename := filepath.Join(directory, "example-git-file")
	err = os.WriteFile(filename, []byte("hello world!"), 0o644)
	CheckIfError(err)

	Info("git add example-git-file")
	_, err = w.Add("example-git-file")
	CheckIfError(err)

	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)

	fmt.Println(status)

	Info("git commit -m \"example go-git commit\"")
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}
