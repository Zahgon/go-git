package main

import (
	"log"
	"os"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
	"github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

func main() {
	CheckArgs("<ssh-url>", "<directory>", "<tag>", "<name>", "<email>", "<public-key>")
	url := os.Args[1]
	directory := os.Args[2]
	tag := os.Args[3]
	key := os.Args[6]

	r, err := cloneRepo(url, directory, key)
	if err != nil {
		log.Printf("clone repo error: %s", err)
		return
	}
	defer func() { _ = r.Close() }()

	created, err := setTag(r, tag)
	if err != nil {
		log.Printf("create tag error: %s", err)
		return
	}

	if created {
		err = pushTags(r, key)
		if err != nil {
			log.Printf("push tag error: %s", err)
			return
		}
	}
}

func cloneRepo(url, dir, publicKeyPath string) (*git.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func publicKey(filePath string) (*ssh.PublicKeys, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tagExists(tag string, r *git.Repository) bool { _ = "STUB: not implemented"; return false }

func setTag(r *git.Repository, tag string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func pushTags(r *git.Repository, publicKeyPath string) error { _ = "STUB: not implemented"; return nil }
