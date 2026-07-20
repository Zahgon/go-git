package http

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	transport "github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/go-git/go-git/v6/storage"
)

func (s *dumbPackSession) fetchDumb(ctx context.Context, st storage.Storer, req *transport.FetchRequest) error {
	_ = "STUB: not implemented"
	return nil
}

type fetchWalker struct {
	ctx        context.Context
	client     *http.Client
	baseURL    *url.URL
	authorizer func(*http.Request) error
	st         storage.Storer
	refs       *packp.AdvRefs
	fs         billy.Filesystem
	queue      []plumbing.Hash
	packIdx    map[plumbing.Hash]string
}

func newFetchWalker(ctx context.Context, s *dumbPackSession, st storage.Storer, fs billy.Filesystem) *fetchWalker {
	_ = "STUB: not implemented"
	return nil
}

func (r *fetchWalker) httpGet(urlPath string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *fetchWalker) getInfoPacks() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *fetchWalker) downloadFile(fp string) (rErr error) { _ = "STUB: not implemented"; return nil }

func (r *fetchWalker) getHead() (ref *plumbing.Reference, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *fetchWalker) process() error { _ = "STUB: not implemented"; return nil }

func (r *fetchWalker) fetchObject(objHash plumbing.Hash, obj plumbing.EncodedObject) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func objectFormatFromHash(h plumbing.Hash) formatcfg.ObjectFormat {
	_ = "STUB: not implemented"
	return *new(formatcfg.ObjectFormat)
}

func (r *fetchWalker) fetch() error { _ = "STUB: not implemented"; return nil }
