package http

import "net/http"

type BasicAuth struct {
	Username, Password string
}

func (a *BasicAuth) Authorizer(r *http.Request) error { _ = "STUB: not implemented"; return nil }

type TokenAuth struct {
	Token string
}

func (a *TokenAuth) Authorizer(r *http.Request) error { _ = "STUB: not implemented"; return nil }
