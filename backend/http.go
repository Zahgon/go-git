package backend

import (
	"net/http"
	"regexp"

	"github.com/go-git/go-git/v6/plumbing/transport"
)

type httpService struct {
	pattern *regexp.Regexp
	method  string
	handler func(b *Backend, w http.ResponseWriter, r *http.Request, repo, file, svc string)
	svc     string
}

var httpServices = []httpService{
	{regexp.MustCompile("(.*?)/HEAD$"), http.MethodGet, (*Backend).handleDumbTextFile, ""},
	{regexp.MustCompile("(.*?)/info/refs$"), http.MethodGet, (*Backend).handleInfoRefs, ""},
	{regexp.MustCompile("(.*?)/objects/info/alternates$"), http.MethodGet, (*Backend).handleDumbTextFile, ""},
	{regexp.MustCompile("(.*?)/objects/info/http-alternates$"), http.MethodGet, (*Backend).handleDumbTextFile, ""},
	{regexp.MustCompile("(.*?)/objects/info/packs$"), http.MethodGet, (*Backend).handleDumbInfoPacks, ""},
	{regexp.MustCompile("(.*?)/objects/[0-9a-f]{2}/[0-9a-f]{38,62}$"), http.MethodGet, (*Backend).handleDumbLooseObject, ""},
	{regexp.MustCompile(`(.*?)/objects/pack/pack-[0-9a-f]{40,64}\.pack$`), http.MethodGet, (*Backend).handleDumbPackFile, ""},
	{regexp.MustCompile(`(.*?)/objects/pack/pack-[0-9a-f]{40,64}\.idx$`), http.MethodGet, (*Backend).handleDumbIdxFile, ""},
	{regexp.MustCompile("(.*?)/git-upload-pack$"), http.MethodPost, (*Backend).handleServiceRPC, transport.UploadPackService},
	{regexp.MustCompile("(.*?)/git-receive-pack$"), http.MethodPost, (*Backend).handleServiceRPC, transport.ReceivePackService},
	{regexp.MustCompile("(.*?)/git-upload-archive$"), http.MethodPost, (*Backend).handleServiceRPC, transport.UploadArchiveService},
}

func (b *Backend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleServiceRPC(w http.ResponseWriter, r *http.Request, repo, _, svc string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleInfoRefs(w http.ResponseWriter, r *http.Request, repo, file, _ string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleDumbTextFile(w http.ResponseWriter, r *http.Request, repo, file, _ string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleDumbInfoPacks(w http.ResponseWriter, r *http.Request, repo, file, _ string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleDumbLooseObject(w http.ResponseWriter, r *http.Request, repo, file, _ string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleDumbPackFile(w http.ResponseWriter, r *http.Request, repo, file, _ string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleDumbIdxFile(w http.ResponseWriter, r *http.Request, repo, file, _ string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) handleDumbSendFile(w http.ResponseWriter, _ *http.Request, repo, file, contentType string) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) requireReceivePackAuth(w http.ResponseWriter, r *http.Request, service string) bool {
	_ = "STUB: not implemented"
	return false
}

func renderStatusError(w http.ResponseWriter, code int) { _ = "STUB: not implemented"; return }

func hdrNocache(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func hdrCacheForever(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
