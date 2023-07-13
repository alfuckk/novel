package endpoint

import (
	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	ossService "github.com/kzaun/novel/internal/oss/service"
	"github.com/kzaun/novel/pkg/middleware"
)

type Endpoints struct {
	FPutObject endpoint.Endpoint
}

func MakeEndpoints(s ossService.Service, logger log.Logger) Endpoints {
	return Endpoints{
		FPutObject: middleware.LoggingMiddleware(logger)(makeFPutObjectEndpoint(s)),
	}
}
