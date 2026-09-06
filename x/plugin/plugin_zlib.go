package plugin

import (
	xzlib "github.com/go-git/go-git/v6/x/plugin/zlib"
)

func init() {

	_ = Register(Zlib(), func() ZlibProvider {
		return xzlib.NewStdlib()
	})
}

const zlibPlugin Name = "zlib"

var zlibKey = newKeyWithValidator(zlibPlugin, xzlib.ValidateProvider)

type ZlibReader = xzlib.Reader

type ZlibWriter = xzlib.Writer

type ZlibProvider = xzlib.Provider

func Zlib() key[ZlibProvider] { _ = "STUB: not implemented"; return nil }
