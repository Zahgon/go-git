package main

import (
	"log"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/storage/memory"
)

func main() {
	CheckArgs("<url>")
	url := os.Args[1]

	Info("git ls-remote --tags %s", url)

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

	log.Print("Fetching tags...")

	refs, err := rem.List(&git.ListOptions{

		PeelingOption: git.AppendPeeled,
	})
	if err != nil {
		log.Fatal(err)
	}

	var tags []string
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
	}

	if len(tags) == 0 {
		log.Println("No tags!")
		return
	}

	log.Printf("Tags found: %v", tags)
}
