package main

import (
	"github.com/kzaun/novel/internal/scraper"
	"github.com/kzaun/novel/pkg/fx/knadhfx"
	"github.com/kzaun/novel/pkg/fx/logfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		knadhfx.ConfigModule,
		logfx.LogModule,
		scraper.ScraperModule,
		fx.Invoke(scraper.Register),
	).Run()
}
