package main

import (
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/go-git/go-git/v6/storage/filesystem"
)

func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	fs := r.Storer.(*filesystem.Storage).Filesystem()
	err = transport.UpdateServerInfo(r.Storer, fs)
	CheckIfError(err)
}
