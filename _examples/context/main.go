package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-git/go-git/v6"
	. "github.com/go-git/go-git/v6/_examples"
)

func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	Info("git clone %s %s", url, directory)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-stop
		Warning("\nSignal detected, canceling operation...")
		cancel()
	}()

	Warning("To gracefully stop the clone operation, push Crtl-C.")

	r, err := git.PlainCloneContext(ctx, directory, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	CheckIfError(err)
	defer func() { _ = r.Close() }()
}
