package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func prepareRepo(w *git.Worktree, directory string) { _ = "STUB: not implemented"; return }

func main() {
	CheckArgs("<directory>")
	directory := os.Args[1]

	r, err := git.PlainOpen(directory)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	w, err := r.Worktree()
	CheckIfError(err)

	prepareRepo(w, directory)

	Info("echo \"hello world! Modify 2\" > for-modify")
	err = os.WriteFile(filepath.Join(directory, "for-modify"), []byte("hello world! Modify 2"), 0o644)
	CheckIfError(err)
	Info("git add for-modify")
	_, err = w.Add("for-modify")
	CheckIfError(err)

	Info("echo \"hello world! Add\" > for-add")
	err = os.WriteFile(filepath.Join(directory, "for-add"), []byte("hello world! Add"), 0o644)
	CheckIfError(err)
	Info("git add for-add")
	_, err = w.Add("for-add")
	CheckIfError(err)

	Info("rm for-delete")
	err = os.Remove(filepath.Join(directory, "for-delete"))
	CheckIfError(err)
	Info("git add for-delete")
	_, err = w.Add("for-delete")
	CheckIfError(err)

	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)
	fmt.Println(status)

	Info("git restore --staged for-modify")
	err = w.Restore(&git.RestoreOptions{Staged: true, Files: []string{"for-modify"}})
	CheckIfError(err)

	Info("git status --porcelain")
	status, err = w.Status()
	CheckIfError(err)
	fmt.Println(status)

	Info("git restore --staged for-add for-delete")
	err = w.Restore(&git.RestoreOptions{Staged: true, Files: []string{"for-add", "for-delete"}})
	CheckIfError(err)

	Info("git status --porcelain")
	status, err = w.Status()
	CheckIfError(err)
	fmt.Println(status)
}
