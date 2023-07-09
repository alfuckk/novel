package ginfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type GinServiceImpl struct {
	engine *gin.Engine
}

// Define your service interface and implementation
type GinService interface {
	// UploadFile(bucketName, objectName, filePath string) error
}


func NewGinClient() (*gin.Engine) {
	return gin.Default()
}

func NewGinService(engine *gin.Engine) GinService {
	return &GinServiceImpl{
		engine: engine,
	}
}

// Define your application module
func NewGinModule() fx.Option {
	return fx.Options(
		fx.Provide(NewGinClient),
		fx.Provide(NewGinService),
	)
}

func (s *GinServiceImpl) UploadFile(bucketName, objectName, filePath string) error {
	// Implement the file upload logic using the MinIO client
	// _,err := s.client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	// if err != nil {
	// 	return err
	// }
	return nil
}