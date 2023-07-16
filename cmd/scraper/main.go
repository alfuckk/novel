package main

import (
	"github.com/kzaun/novel/internal/scraper"
	"github.com/kzaun/novel/pkg/fx/knadhfx"
	"github.com/kzaun/novel/pkg/fx/logfx"
	"github.com/kzaun/novel/pkg/fx/rabbitmqfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		knadhfx.ConfigModule,
		logfx.LogModule,
		scraper.ScraperModule,
		rabbitmqfx.MQModule,
		fx.Invoke(scraper.Register),
	).Run()
}
