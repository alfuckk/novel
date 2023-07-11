package oss

import (
	"context"
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
