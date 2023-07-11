package oss

import (
	"go.uber.org/fx"
)

var Module = fx.Module("oss",
	fx.Provide(
		OssModule,
		// lib.AsRoute(handler.ProvideEchoHandler),
		// lib.AsRoute(handler.ProvideUploadHandler),
	),
)
