package service

import (
	"context"
	"path"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/knadh/koanf/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/fx"
)

type Service interface {
	Fput(ctx context.Context, bucketName, tmpPath, contentType string) (imageURL string, err error)
}
type Params struct {
	fx.In
	Config *koanf.Koanf
	Logger log.Logger
}
type service struct {
	Client     *minio.Client
	OssGateway string
	Log        log.Logger
}

func (s service) Fput(ctx context.Context, bucketName, tmpPath, contentType string) (imageURL string, err error) {
	filesuffix := path.Ext(tmpPath)
	objectName := strconv.Itoa(int(time.Now().Unix())) + filesuffix
	_, err = s.Client.FPutObject(ctx, bucketName, objectName, tmpPath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		s.Log.Log("upload %s failed: %v", err)
		return "", err
	}
	imageURL = "https://" + s.OssGateway + "/" + bucketName + "/" + objectName
	return imageURL, nil
}

func NewService(p Params) Service {
	client, err := minio.New(p.Config.String("OSS.Endpoint"), &minio.Options{
		Creds:  credentials.NewStaticV4(p.Config.String("OSS.AccessKey"), p.Config.String("OSS.SecretKey"), ""),
		Secure: p.Config.Bool("OSS.UseSSL"),
	})
	if err != nil {
		p.Logger.Log("minio sdk init faild.")
		return nil
	}
	p.Logger.Log("minio client init success.")
	return &service{
		Client:     client,
		OssGateway: p.Config.String("OSS.Endpoint"),
		Log:        p.Logger,
	}
}
