package main

import (
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/config"
)

func main() {
	tmp, err := os.MkdirTemp("", "go-git-example")
	CheckIfError(err)
	defer os.RemoveAll(tmp)

	Info("git init")
	r, err := git.PlainInit(tmp, false)
	CheckIfError(err)
	defer func() { _ = r.Close() }()

	cfg, err := r.Config()
	CheckIfError(err)

	Info("worktree is %s", cfg.Core.Worktree)

	cfg.Remotes["origin"] = &config.RemoteConfig{
		Name: "origin",
		URLs: []string{"https://github.com/git-fixtures/basic.git"},
	}

	Info("origin remote: %+v", cfg.Remotes["origin"])

	cfg.User.Name = "Local name"

	Info("custom.name is %s", cfg.User.Name)

	r.Storer.SetConfig(cfg)
}
