//go:build !darwin && !linux

package mmap

import (
	"github.com/go-git/go-billy/v6"
)

type PackScanner struct{}

func NewPackScanner(_ int, _, _, _ billy.File) (*PackScanner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
