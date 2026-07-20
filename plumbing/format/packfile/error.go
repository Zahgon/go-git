package packfile

type Error struct {
	reason, details string
}

func NewError(reason string) *Error { _ = "STUB: not implemented"; return nil }

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) AddDetails(format string, args ...any) *Error {
	_ = "STUB: not implemented"
	return nil
}
