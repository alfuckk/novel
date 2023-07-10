package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		fx.Invoke(register),
	).Run()
}

func register(lc fx.Lifecycle, router *gin.Engine) {
	endpoints := makeEndpoints()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				router.GET("/endpoint", gin.WrapF(httptransport.NewServer(
					endpoints.Endpoint,
					decodeRequest,
					encodeResponse,
				)))
				_ = http.ListenAndServe(":8080", router)
			}()
			return nil
		},
	})
}

func makeEndpoints() Endpoints {
	return Endpoints{
		Endpoint: endpoint.Endpoint(func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// Your business logic here
			return nil, nil
		}),
	}
}

type Endpoints struct {
	Endpoint endpoint.Endpoint
}