package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing"
)

func main() {
	CheckArgs("<url>", "<directory>", "<branch>")
	url, directory, branch := os.Args[1], os.Args[2], os.Args[3]

	Info("git clone %s %s", url, directory)
	r, err := git.PlainClone(directory, &git.CloneOptions{
		URL: url,
	})
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git show-ref --head HEAD")
	ref, err := r.Head()
	CheckIfError(err)

	fmt.Println(ref.Hash())

	w, err := r.Worktree()
	CheckIfError(err)

	Info("git checkout %s", branch)

	branchRefName := plumbing.NewBranchReferenceName(branch)
	branchCoOpts := git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchRefName),
		Force:  true,
	}
	if err := w.Checkout(&branchCoOpts); err != nil {
		Warning("local checkout of branch '%s' failed, will attempt to fetch remote branch of same name.", branch)
		Warning("like `git checkout <branch>` defaulting to `git checkout -b <branch> --track <remote>/<branch>`")

		mirrorRemoteBranchRefSpec := fmt.Sprintf("refs/heads/%s:refs/heads/%s", branch, branch)
		err = fetchOrigin(r, mirrorRemoteBranchRefSpec)
		CheckIfError(err)

		err = w.Checkout(&branchCoOpts)
		CheckIfError(err)
	}
	CheckIfError(err)

	Info("checked out branch: %s", branch)

	Info("git show-ref --head HEAD")
	ref, err = r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())
}

func fetchOrigin(repo *git.Repository, refSpecStr string) error {
	_ = "STUB: not implemented"
	return nil
}
