package main

import (
	"github.com/go-git/go-git/v6/plumbing/object"
)

func checkIfError(err error, code exitCode, mainReason string, v ...any) {
	_ = "STUB: not implemented"
	return
}

func helpAndExit(s, helpMsg string, code exitCode) { _ = "STUB: not implemented"; return }

func printErr(err error) { _ = "STUB: not implemented"; return }

func printMsg(format string, args ...any) { _ = "STUB: not implemented"; return }

func printCommits(commits []*object.Commit) { _ = "STUB: not implemented"; return }

func wrapErr(err error, s string, v ...any) error { _ = "STUB: not implemented"; return nil }
