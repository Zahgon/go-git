package main

import (
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func main() {
	CheckArgs("<repository-path>")
	path := os.Args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git push")

	err = r.Push(&git.PushOptions{})
	CheckIfError(err)
}
