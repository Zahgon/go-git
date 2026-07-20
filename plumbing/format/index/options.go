package index

type Option func(*options)

type options struct {
	skipHash bool
}

func WithSkipHash() Option { _ = "STUB: not implemented"; return *new(Option) }
