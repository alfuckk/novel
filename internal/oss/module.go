package oss

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/knadh/koanf/v2"
	ossEndpoint "github.com/kzaun/novel/internal/oss/endpoint"
	ossService "github.com/kzaun/novel/internal/oss/service"
	ossTransport "github.com/kzaun/novel/internal/oss/transport"
	"go.uber.org/fx"
)

var OssModule = fx.Module("oss",
	fx.Provide(
		ossService.NewService,
		ossEndpoint.MakeEndpoints,
		ossTransport.NewHTTPHandler,
	),
)

func Register(
	lc fx.Lifecycle,
	handler http.Handler,
	logger log.Logger,
	config *koanf.Koanf,
) {
	errs := make(chan error)
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				c := make(chan os.Signal, 1)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				errs <- fmt.Errorf("%s", <-c)
			}()
			go func() {
				logger.Log("transport", "HTTP", "addr", config.String("Port"))
				errs <- http.ListenAndServe(config.String("Port"), handler)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 停止服务的逻辑
			logger.Log("exit", <-errs)
			return nil
		},
	})
}
