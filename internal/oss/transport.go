package oss

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
)

func Register(lc fx.Lifecycle, router *gin.Engine, s Service, p *koanf.Koanf) {
	endpoints := makeEndpoints(s)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				server := httptransport.NewServer(
					endpoints.GetUserEndpoint,
					decodeGetUserRequest,
					encodeResponse,
				)
				router.GET("/get-user", func(c *gin.Context) {
					server.ServeHTTP(c.Writer, c.Request)
				})
				_ = http.ListenAndServe(p.String("oss.port"), router)
			}()
			return nil
		},
	})
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
