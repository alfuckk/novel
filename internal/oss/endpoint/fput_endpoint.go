package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kzaun/novel/internal/oss/service"
)

func makeFPutObjectEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FPutObjectRequest)

		v, _ := s.Fput(ctx, req.BucketName, req.FilePath, req.ContentType)
		return FPutObjectResponse{OssURL: v}, nil
	}
}

type FPutObjectRequest struct {
	BucketName  string
	FilePath    string
	FileName    string
	ContentType string
}

type FPutObjectResponse struct {
	OssURL string `json:"ossURL"`
}
