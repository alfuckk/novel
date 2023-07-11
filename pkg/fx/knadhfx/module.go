package knadhfx

import (
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
)

func ProvideConfig() (*koanf.Koanf, error) {
	k := koanf.New(".")
	if _, exists := os.LookupEnv("debug"); !exists {
		err := k.Load(file.Provider("app.yaml"), yaml.Parser())
		if err != nil {
			return nil, err
		}
	}
	return k, nil
}

// ConfigModule provided to fx
var ConfigModule = fx.Options(
	fx.Provide(ProvideConfig),
)
