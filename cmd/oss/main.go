package main

import (
	"github.com/kzaun/novel/internal/oss"
	"github.com/kzaun/novel/pkg/fx/knadhfx"
	"github.com/kzaun/novel/pkg/fx/logfx"
	"github.com/kzaun/novel/pkg/fx/miniofx"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		miniofx.MinioModule,
		knadhfx.ConfigModule,
		logfx.LogModule,
		oss.OssModule,
		fx.Invoke(oss.Register),
	).Run()
}
