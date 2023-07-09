package logfx

import (
	"os"

	"github.com/go-kit/log"
)

func ProvideLogger() log.Logger {
	return log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
}

