package miniofx

import (
	"github.com/minio/minio-go/v7"
	"go.uber.org/fx"
)

// ProvideMinio is a function that returns a minio client instance
func ProvideMinio() (*minio.Client, error) {
	// TODO: Initialize and return a minio client
	return nil, nil
}

// Module provided to fx
var MinioModule = fx.Options(
	fx.Provide(ProvideMinio),
)
