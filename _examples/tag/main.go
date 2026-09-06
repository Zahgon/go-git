package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
)

func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	Info("git show-ref --tag")

	tagrefs, err := r.Tags()
	CheckIfError(err)
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		fmt.Println(t)
		return nil
	})
	CheckIfError(err)

	Info("for t in $(git show-ref --tag); do if [ \"$(git cat-file -t $t)\" = \"tag\" ]; then git cat-file -p $t ; fi; done")

	tags, err := r.TagObjects()
	CheckIfError(err)
	err = tags.ForEach(func(t *object.Tag) error {
		fmt.Println(t)
		return nil
	})
	CheckIfError(err)
}
