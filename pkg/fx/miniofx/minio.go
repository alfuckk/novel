package miniofx

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type Result struct {
	fx.Out

	Client *minio.Client
}

func New(p Params) (Result, error) {
	client, err := minio.New(p.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(p.AccessKey, p.SecretKey, ""),
		Secure: p.UseSSL,
	})
	if err != nil {
		return Result{}, err
	}

	return Result{Client: client}, nil
}