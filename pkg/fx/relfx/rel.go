package relfx

import (
	"go.uber.org/fx"
)

type Module struct {
	fx.In
}

func New() *Module {
	return &Module{}
}
