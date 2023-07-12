package logfx

import (
	"os"

	kitlog "github.com/go-kit/log"

	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Logger kitlog.Logger
}

var LogModule = fx.Provide(func() kitlog.Logger {
	return kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
})
