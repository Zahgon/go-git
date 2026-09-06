package main

import (
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	Info("git clone %s %s", url, directory)

	r, err := git.PlainClone(directory, &git.CloneOptions{
		URL:   url,
		Depth: 1,

		Progress: os.Stdout,
	})

	CheckIfError(err)
	defer func() { _ = r.Close() }()
}
