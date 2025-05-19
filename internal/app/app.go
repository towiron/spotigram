package app

import (
	"github.com/towiron/spotigram/internal/adapter"
	repository "github.com/towiron/spotigram/internal/respository"
	"github.com/towiron/spotigram/internal/service"
	"go.uber.org/fx"
)

func New(opt fx.Option) *fx.App {
	return fx.New(
		opt,
		adapter.Module,
		service.Module,
		repository.Module,
	)
}
