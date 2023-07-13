package transport

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/kzaun/novel/internal/oss/endpoint"
)

func NewHTTPHandler(endpoints endpoint.Endpoints, logger log.Logger) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/fput", makeFputHandler(endpoints, logger))
	return r
}

func makeFputHandler(endpoints endpoint.Endpoints, logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取文件
		file, _ := c.FormFile("file")
		// 获取表单字段
		field := c.PostForm("bucket")

		// 对表单字段进行校验
		if field == "" {
			c.String(http.StatusBadRequest, "字段值不能为空")
			return
		}

		// 保存文件到系统的临时目录
		tmpDir := os.TempDir()
		dst := filepath.Join(tmpDir, file.Filename)
		c.SaveUploadedFile(file, dst)

		// 读取文件的一部分内容
		f, _ := os.Open(dst)
		defer f.Close()
		buffer := make([]byte, 512)
		f.Read(buffer)

		// 获取content-type
		contentType := http.DetectContentType(buffer)

		request := endpoint.FPutObjectRequest{FilePath: dst, FileName: file.Filename, ContentType: contentType, BucketName: field}
		response, err := endpoints.FPutObject(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
