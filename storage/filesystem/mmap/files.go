//go:build darwin || linux

package mmap

import (
	"github.com/go-git/go-billy/v6"
)

func mmapFile(f billy.File) ([]byte, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getFileDescriptor(f billy.File) (uintptr, error) { _ = "STUB: not implemented"; return 0, nil }

func validateFile(mmap []byte, sv uint32, sig []byte, minLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type billyFileDescriptor interface {
	Fd() (uintptr, bool)
}

type goFileDescriptor interface {
	Fd() uintptr
}
