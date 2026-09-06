package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/storage/memory"
)

func main() {

	Info("git clone https://github.com/src-d/go-siva")
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/src-d/go-siva",
	})
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git log")

	ref, err := r.Head()
	CheckIfError(err)

	since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2019, 7, 30, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	CheckIfError(err)

	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	CheckIfError(err)
}
