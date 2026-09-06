//go:build windows

package sshagent

import (
	"errors"
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

const MaxMessageLen = 8192

var (
	ErrPageantNotFound = errors.New("pageant process not found")

	ErrSendMessage = errors.New("error sending message")

	ErrMessageTooLong = errors.New("message too long")

	ErrInvalidMessageFormat = errors.New("invalid message format")

	ErrResponseTooLong = errors.New("response too long")
)

const (
	agentCopydataID = 0x804e50ba
	wmCopydata      = 74
)

type copyData struct {
	dwData uintptr
	cbData uint32
	lpData unsafe.Pointer
}

var (
	lock sync.Mutex

	user32dll      = windows.NewLazySystemDLL("user32.dll")
	winFindWindow  = winAPI(user32dll, "FindWindowW")
	winSendMessage = winAPI(user32dll, "SendMessageW")

	kernel32dll           = windows.NewLazySystemDLL("kernel32.dll")
	winGetCurrentThreadID = winAPI(kernel32dll, "GetCurrentThreadId")
)

func winAPI(dll *windows.LazyDLL, funcName string) func(...uintptr) (uintptr, uintptr, error) {
	_ = "STUB: not implemented"
	return nil
}

func query(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func pageantWindow() uintptr { _ = "STUB: not implemented"; return 0 }
