package main

import (
	"context"
	"log"

	"github.com/kzaun/novel/internal/oss"
	"github.com/kzaun/novel/pkg/fx/ginfx"
	"github.com/kzaun/novel/pkg/fx/miniofx"

	// "github.com/go-kit/kit/transport/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		miniofx.NewMinIOModule(),
		ginfx.NewGinModule(),
		fx.Invoke(oss.RegisterHandlers),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}


