package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-git/go-billy/v6"
	"github.com/go-git/go-billy/v6/osfs"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/object/commitgraph"
	"github.com/go-git/go-git/v6/storage/filesystem"
)

func main() {
	CheckArgs("<path>", "<revision>", "<tree path>")

	path := os.Args[1]
	revision := os.Args[2]
	treePath := os.Args[3]

	fs := osfs.New(path)
	if _, err := fs.Stat(git.GitDirName); err == nil {
		fs, err = fs.Chroot(git.GitDirName)
		CheckIfError(err)
	}

	s := filesystem.NewStorage(fs, cache.NewObjectLRUDefault())
	r, err := git.Open(s, fs)
	CheckIfError(err)
	defer s.Close()

	Info("git rev-parse %s", revision)

	h, err := r.ResolveRevision(plumbing.Revision(revision))
	CheckIfError(err)

	commit, err := r.CommitObject(*h)
	CheckIfError(err)

	tree, err := commit.Tree()
	CheckIfError(err)
	if treePath != "" {
		tree, err = tree.Tree(treePath)
		CheckIfError(err)
	}

	var paths []string
	for _, entry := range tree.Entries {
		paths = append(paths, entry.Name)
	}

	commitNodeIndex, file := getCommitNodeIndex(r, fs)
	if file != nil {
		defer file.Close()
	}

	commitNode, err := commitNodeIndex.Get(*h)
	CheckIfError(err)

	revs, err := getLastCommitForPaths(commitNode, treePath, paths)
	CheckIfError(err)
	for path, rev := range revs {

		hash := rev.Hash.String()
		line := strings.Split(rev.Message, "\n")
		fmt.Println(path, hash[:7], line[0])
	}
}

func getCommitNodeIndex(r *git.Repository, fs billy.Filesystem) (commitgraph.CommitNodeIndex, io.ReadCloser) {
	_ = "STUB: not implemented"
	return *new(commitgraph.CommitNodeIndex), *new(io.ReadCloser)
}

type commitAndPaths struct {
	commit commitgraph.CommitNode

	paths []string

	hashes map[string]plumbing.Hash
}

func getCommitTree(c commitgraph.CommitNode, treePath string) (*object.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFullPath(treePath, path string) string { _ = "STUB: not implemented"; return "" }

func getFileHashes(c commitgraph.CommitNode, treePath string, paths []string) (map[string]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getLastCommitForPaths(c commitgraph.CommitNode, treePath string, paths []string) (map[string]*object.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
