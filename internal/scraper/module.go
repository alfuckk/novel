package scraper

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/knadh/koanf/v2"
	"github.com/kzaun/novel/internal/scraper/service"

	"go.uber.org/fx"
)

var ScraperModule = fx.Module("scraper",
	fx.Provide(
		service.NewService,
	),
)

func Register(
	lc fx.Lifecycle,
	logger log.Logger,
	s service.Service,
	config *koanf.Koanf,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				result, err := s.Scrape("https://quanben.io")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 停止服务的逻辑
			return nil
		},
	})
}
