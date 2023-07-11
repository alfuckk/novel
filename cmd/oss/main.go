package main

import (
	"context"
	"log"

	"github.com/kzaun/novel/internal/oss"
	"github.com/kzaun/novel/pkg/fx/ginfx"
	"github.com/kzaun/novel/pkg/fx/knadhfx"
	"github.com/kzaun/novel/pkg/fx/logfx"
	"github.com/kzaun/novel/pkg/fx/miniofx"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		miniofx.MinioModule,
		ginfx.GinModule,
		oss.OssModule,
		knadhfx.ConfigModule,
		logfx.LogModule,
		fx.Invoke(oss.Register),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
