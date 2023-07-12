package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kzaun/novel/internal/oss/service"
)

func makeFPutObjectEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FPutObjectRequest)
		v := s.Echo(req.S)
		return FPutObjectResponse{V: v}, nil
	}
}

type FPutObjectRequest struct {
	S string
}

type FPutObjectResponse struct {
	V string `json:"version"`
}
