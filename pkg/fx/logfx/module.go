package logfx

import (
	"os"

	"github.com/go-kit/log"

	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Logger log.Logger
}

var LogModule = fx.Provide(func() log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	return logger
})
