package plugin

import (
	"github.com/go-git/go-git/v6/config"
	xconfig "github.com/go-git/go-git/v6/x/plugin/config"
)

func init() {

	_ = Register(ConfigLoader(), func() ConfigSource {
		return xconfig.NewAuto()
	})
}

const configLoaderPlugin Name = "config-loader"

var configLoader = newKey[ConfigSource](configLoaderPlugin)

type ConfigSource interface {
	Load(scope config.Scope) (config.ConfigStorer, error)
}

func ConfigLoader() key[ConfigSource] {
	_ = "STUB: not implemented" //nolint:revive // intentional unexported return type
	return nil
}
