package memory

import formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"

type options struct {
	objectFormat formatcfg.ObjectFormat
}

func newOptions() options { _ = "STUB: not implemented"; return *new(options) }

type StorageOption func(*options)

func WithObjectFormat(of formatcfg.ObjectFormat) StorageOption {
	_ = "STUB: not implemented"
	return *new(StorageOption)
}
