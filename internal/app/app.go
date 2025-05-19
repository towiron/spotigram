package app

import "go.uber.org/fx"

func New(opt fx.Option) *fx.App {
	return fx.New(
		opt,
	)
}
