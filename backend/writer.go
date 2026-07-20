package backend

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

const defaultChunkSize = 4096

type flushResponseWriter struct {
	http.ResponseWriter
	log       *log.Logger
	chunkSize int

	started atomic.Bool
}

func (f *flushResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (f *flushResponseWriter) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *flushResponseWriter) ReadFrom(r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *flushResponseWriter) Close() error { _ = "STUB: not implemented"; return nil }
