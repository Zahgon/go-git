package config

type Option struct {
	Key string

	Value string
}

type Options []*Option

func (o *Option) IsKey(key string) bool { _ = "STUB: not implemented"; return false }

func (opts Options) GoString() string { _ = "STUB: not implemented"; return "" }

func (opts Options) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (opts Options) Has(key string) bool { _ = "STUB: not implemented"; return false }

func (opts Options) GetAll(key string) []string { _ = "STUB: not implemented"; return nil }

func (opts Options) withoutOption(key string) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func (opts Options) withAddedOption(key, value string) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func (opts Options) withSettedOption(key string, values ...string) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}
