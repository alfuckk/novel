package oss

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log"
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
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				http.ListenAndServe(":8080", handler)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 停止服务的逻辑
			return nil
		},
	})
}
