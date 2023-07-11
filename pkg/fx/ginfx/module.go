package ginfx

import (
	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

// ProvideGin 是一个返回 gin.Engine 实例的函数
func ProvideGin() *gin.Engine {
	r := gin.New()
	return r
}

// Module 是提供给 fx 的模块
var GinModule = fx.Options(
	fx.Provide(ProvideGin),
)
