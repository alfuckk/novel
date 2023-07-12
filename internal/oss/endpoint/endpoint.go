package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kzaun/novel/internal/oss/service"
)

type Endpoints struct {
	Echo endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Echo: makeEchoEndpoint(s),
	}
}

func makeEchoEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EchoRequest)
		v := s.Echo(req.S)
		return EchoResponse{V: v}, nil
	}
}

type EchoRequest struct {
	S string
}

type EchoResponse struct {
	V string `json:"version"`
}
