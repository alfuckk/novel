package oss

import (
	"go.uber.org/fx"
)

var OssModule = fx.Module("oss",
	fx.Provide(
		NewService,
		// lib.AsRoute(handler.ProvideEchoHandler),
		// lib.AsRoute(handler.ProvideUploadHandler),
	),
)
