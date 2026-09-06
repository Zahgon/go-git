package noder

import "fmt"

type Hasher interface {
	Hash() []byte
}

type Equal func(a, b Hasher) bool

type Noder interface {
	Hasher
	fmt.Stringer

	Name() string

	IsDir() bool

	Children() ([]Noder, error)

	NumChildren() (int, error)
	Skip() bool
}

var NoChildren = []Noder{}
