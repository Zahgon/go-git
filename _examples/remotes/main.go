package main

import (
	"fmt"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/storage/memory"
)

func main() {

	Info("git init")
	r, err := git.Init(memory.NewStorage())
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git remote add example https://github.com/git-fixtures/basic.git")
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "example",
		URLs: []string{"https://github.com/git-fixtures/basic.git"},
	})

	CheckIfError(err)

	Info("git remote -v")

	list, err := r.Remotes()
	CheckIfError(err)

	for _, r := range list {
		fmt.Println(r)
	}

	Info("git fetch example")
	err = r.Fetch(&git.FetchOptions{
		RemoteName: "example",
	})

	CheckIfError(err)

	Info("git show-ref")

	refs, err := r.References()
	CheckIfError(err)

	err = refs.ForEach(func(ref *plumbing.Reference) error {

		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}

		fmt.Println(ref)
		return nil
	})

	CheckIfError(err)

	Info("git remote rm example")

	err = r.DeleteRemote("example")
	CheckIfError(err)
}
