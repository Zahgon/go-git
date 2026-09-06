package trace

import (
	"log"
	"sync/atomic"
)

var (
	logger = newLogger()

	current atomic.Int32
)

func newLogger() *log.Logger { _ = "STUB: not implemented"; return nil }

type Target int32

const (
	General Target = 1 << iota

	Packet

	SSH

	Performance

	HTTP

	Internal
)

func SetTarget(target Target) { _ = "STUB: not implemented"; return }

func SetLogger(l *log.Logger) { _ = "STUB: not implemented"; return }

func (t Target) Print(args ...any) { _ = "STUB: not implemented"; return }

func (t Target) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (t Target) Enabled() bool { _ = "STUB: not implemented"; return false }

func GetTarget() Target { _ = "STUB: not implemented"; return *new(Target) }
