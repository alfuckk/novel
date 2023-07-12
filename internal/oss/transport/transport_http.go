package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/kzaun/novel/internal/oss/endpoint"
)

func NewHTTPHandler(endpoints endpoint.Endpoints, logger log.Logger) http.Handler {
	r := gin.Default()
	r.GET("/echo", makeEchoHandler(endpoints, logger))
	return r
}

func makeEchoHandler(endpoints endpoint.Endpoints, logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := endpoint.EchoRequest{S: c.Query("s")}
		response, err := endpoints.Echo(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
