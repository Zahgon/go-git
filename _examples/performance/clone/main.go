package main

import (
	"crypto"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/go-git/go-billy/v6/osfs"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/plumbing/hash"
	"github.com/go-git/go-git/v6/storage/filesystem"
)

func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	hash.RegisterHash(crypto.SHA1, sha1.New)

	Info("Set GIT_TRACE_PERFORMANCE=true for a time per operation breakdown")
	Info("git clone --no-tags --depth 1 --single-branch %s %s", url, directory)

	fs := osfs.New(directory)
	dotgit, err := fs.Chroot(".git")
	CheckIfError(err)

	storer := filesystem.NewStorageWithOptions(dotgit, cache.NewObjectLRUDefault(), filesystem.Options{

		HighMemoryMode: true,
	})

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: url,

		Tags: git.NoTags,

		Depth: 1,

		SingleBranch: true,

		Progress: os.Stdout,
	})

	CheckIfError(err)
	defer func() { _ = r.Close() }()

	ref, err := r.Head()
	CheckIfError(err)
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
