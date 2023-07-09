package miniofx

import (
	"context"

	"github.com/minio/minio-go/v7"
	"go.uber.org/fx"
)

type MinIOServiceImpl struct {
	client *minio.Client
}

// Define your service interface and implementation
type MinIOService interface {
	UploadFile(bucketName, objectName, filePath string) error
}


func NewMinIOClient() (*minio.Client, error) {
	// Create and configure the MinIO client
	client, err := minio.New("minio.example.com", &minio.Options{
		Region: "",
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMinIOService(client *minio.Client) MinIOService {
	return &MinIOServiceImpl{
		client: client,
	}
}

// Define your application module
func NewMinIOModule() fx.Option {
	return fx.Options(
		fx.Provide(NewMinIOClient),
		fx.Provide(NewMinIOService),
	)
}

func (s *MinIOServiceImpl) UploadFile(bucketName, objectName, filePath string) error {
	// Implement the file upload logic using the MinIO client
	_,err := s.client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}