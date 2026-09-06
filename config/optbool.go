package config

type OptBool byte

const (
	OptBoolUnset OptBool = iota

	OptBoolFalse

	OptBoolTrue
)

func NewOptBool(v bool) OptBool { _ = "STUB: not implemented"; return *new(OptBool) }

func (o OptBool) IsTrue() bool { _ = "STUB: not implemented"; return false }

func (o OptBool) IsSet() bool { _ = "STUB: not implemented"; return false }

func (o OptBool) String() string { _ = "STUB: not implemented"; return "" }

func (o OptBool) FormatBool() string { _ = "STUB: not implemented"; return "" }

func parseConfigBool(v string) OptBool { _ = "STUB: not implemented"; return *new(OptBool) }
