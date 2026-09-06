package plumbing

type PermanentError struct {
	Err error
}

func NewPermanentError(err error) *PermanentError { _ = "STUB: not implemented"; return nil }

func (e *PermanentError) Error() string { _ = "STUB: not implemented"; return "" }
