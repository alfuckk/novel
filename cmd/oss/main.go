package main

import (
	"github.com/kzaun/novel/internal/oss"
	"github.com/kzaun/novel/pkg/fx/knadhfx"
	"github.com/kzaun/novel/pkg/fx/logfx"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		knadhfx.ConfigModule,
		logfx.LogModule,
		oss.OssModule,
		fx.Invoke(oss.Register),
	).Run()
}
