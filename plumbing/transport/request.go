package transport

import (
	"net/url"

	"github.com/go-git/go-git/v6/plumbing/protocol"
)

type Request struct {
	URL *url.URL

	Command string
	Args    []string

	Protocol protocol.Version
}
