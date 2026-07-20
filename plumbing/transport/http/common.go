package http

import (
	"net/http"
	"net/url"
)

type Err struct {
	URL    *url.URL
	Status int
	Reason string
}

func (e *Err) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (e *Err) Error() string { _ = "STUB: not implemented"; return "" }

func checkError(r *http.Response) error { _ = "STUB: not implemented"; return nil }

const infoRefsPath = "/info/refs"

func applyRedirect(resp *http.Response, baseURL *url.URL) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var safeHeaders = map[string]struct{}{
	"User-Agent":        {},
	"Host":              {},
	"Accept":            {},
	"Content-Type":      {},
	"Content-Length":    {},
	"Cache-Control":     {},
	"Git-Protocol":      {},
	"Transfer-Encoding": {},
	"Content-Encoding":  {},
}

func filterHeaders(h http.Header) http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func redactedURL(u *url.URL) string { _ = "STUB: not implemented"; return "" }

func doRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyAuth(httpReq *http.Request, baseURL *url.URL, authorizer func(*http.Request) error) error {
	_ = "STUB: not implemented"
	return nil
}
