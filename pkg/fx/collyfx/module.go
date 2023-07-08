package collyfx

import "go.uber.org/fx"

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ProvideColly),
)
