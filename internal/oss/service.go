package oss

import (
	"context"

	"go.uber.org/fx"
)

type Service interface {
	GetUser(ctx context.Context, id string) (string, error)
}

type service struct {
	// Your service dependencies here
}

func (s *service) GetUser(ctx context.Context, id string) (string, error) {
	return "User Name", nil
}

func NewService() Service {
	return &service{}
}

var OssModule = fx.Options(
	fx.Provide(NewService), // NewService is a function that returns an instance of Service
)
