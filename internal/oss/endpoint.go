package oss

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetUserEndpoint endpoint.Endpoint
}

func makeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetUserEndpoint: makeGetUserEndpoint(s),
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		name, err := s.GetUser(ctx, req.Id)
		if err != nil {
			return GetUserResponse{Err: err.Error()}, nil
		}
		return GetUserResponse{Name: name}, nil
	}
}

type GetUserRequest struct {
	Id string `json:"id"`
}

type GetUserResponse struct {
	Name string
	Err  string
}
