package gitattributes

import (
	"errors"
	"io"
)

const (
	commentPrefix = "#"
	eol           = "\n"
	macroPrefix   = "[attr]"
)

var (
	ErrMacroNotAllowed = errors.New("macro not allowed")

	ErrInvalidAttributeName = errors.New("invalid attribute name")
)

type MatchAttribute struct {
	Name       string
	Pattern    Pattern
	Attributes []Attribute
}

type attributeState byte

const (
	attributeUnknown     attributeState = 0
	attributeSet         attributeState = 1
	attributeUnspecified attributeState = '!'
	attributeUnset       attributeState = '-'
	attributeSetValue    attributeState = '='
)

type Attribute interface {
	Name() string
	IsSet() bool
	IsUnset() bool
	IsUnspecified() bool
	IsValueSet() bool
	Value() string
	String() string
}

type attribute struct {
	name  string
	state attributeState
	value string
}

func (a attribute) Name() string { _ = "STUB: not implemented"; return "" }

func (a attribute) IsSet() bool { _ = "STUB: not implemented"; return false }

func (a attribute) IsUnset() bool { _ = "STUB: not implemented"; return false }

func (a attribute) IsUnspecified() bool { _ = "STUB: not implemented"; return false }

func (a attribute) IsValueSet() bool { _ = "STUB: not implemented"; return false }

func (a attribute) Value() string { _ = "STUB: not implemented"; return "" }

func (a attribute) String() string { _ = "STUB: not implemented"; return "" }

func ReadAttributes(r io.Reader, domain []string, allowMacro bool) (attributes []MatchAttribute, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseAttributesLine(line string, domain []string, allowMacro bool) (m MatchAttribute, err error) {
	_ = "STUB: not implemented"
	return *new(MatchAttribute), nil
}

func checkMacro(name string, allowMacro bool) (macro bool, macroName string, err error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

func validAttributeName(name string) bool { _ = "STUB: not implemented"; return false }

func unquote(str string) (string, string) { _ = "STUB: not implemented"; return "", "" }
