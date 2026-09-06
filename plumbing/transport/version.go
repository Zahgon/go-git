package transport

import (
	"github.com/go-git/go-git/v6/plumbing/protocol"
	"github.com/go-git/go-git/v6/utils/ioutil"
)

func DiscoverVersion(r ioutil.ReadPeeker) (protocol.Version, error) {
	_ = "STUB: not implemented"
	return *new(protocol.Version), nil
}

func ProtocolVersion(p string) protocol.Version {
	_ = "STUB: not implemented"
	return *new(protocol.Version)
}
