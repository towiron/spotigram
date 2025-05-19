package app

import (
	"github.com/towiron/spotigram/internal/adapter"
	"github.com/towiron/spotigram/internal/service"
	"go.uber.org/fx"
)

func New(opt fx.Option) *fx.App {
	return fx.New(
		opt,
		service.Module,
		adapter.Module,
	)
}
