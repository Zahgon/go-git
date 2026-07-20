package examples

import (
	"github.com/go-git/go-git/v6/internal/trace"
)

func init() {
	trace.ReadEnv()
}

func CheckArgs(arg ...string) { _ = "STUB: not implemented"; return }

func CheckIfError(err error) { _ = "STUB: not implemented"; return }

func Info(format string, args ...any) { _ = "STUB: not implemented"; return }

func Warning(format string, args ...any) { _ = "STUB: not implemented"; return }
