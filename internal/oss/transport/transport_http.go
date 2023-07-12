package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/kzaun/novel/internal/oss/endpoint"
)

func NewHTTPHandler(endpoints endpoint.Endpoints, logger log.Logger) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/fput", makeFputHandler(endpoints, logger))
	return r
}

func makeFputHandler(endpoints endpoint.Endpoints, logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := endpoint.EchoRequest{S: c.Query("s")}
		response, err := endpoints.FPutObject(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
