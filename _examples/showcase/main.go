package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing/object"
)

func main() {
	CheckArgs("<url> <path>")
	url := os.Args[1]
	path := os.Args[2]

	Info("git clone %s %s", url, path)

	r, err := git.PlainClone(path, &git.CloneOptions{URL: url})
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git log -1")

	ref, err := r.Head()
	CheckIfError(err)

	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)
	fmt.Println(commit)

	Info("git ls-tree -r HEAD")

	tree, err := commit.Tree()
	CheckIfError(err)

	tree.Files().ForEach(func(f *object.File) error {
		fmt.Printf("100644 blob %s    %s\n", f.Hash, f.Name)
		return nil
	})

	Info("git log --oneline")

	commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
	CheckIfError(err)

	err = commitIter.ForEach(func(c *object.Commit) error {
		hash := c.Hash.String()
		line := strings.Split(c.Message, "\n")
		fmt.Println(hash[:7], line[0])

		return nil
	})
	CheckIfError(err)
}
