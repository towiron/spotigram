package syncer

import (
	"github.com/towiron/spotigram/internal/interfaces"
	"github.com/towiron/spotigram/internal/pkg/config"
	"github.com/towiron/spotigram/internal/pkg/logger"
	"go.uber.org/fx"
)

type Service struct {
	config config.Configer
	logger logger.Logger
}

type Options struct {
	fx.In
	Logger logger.Logger
	Config config.Configer
}

var Module = fx.Provide(
	fx.Annotate(New, fx.As(new(interfaces.ServiceSyncer))),
)

func New(opts Options) *Service {
	return &Service{
		config: opts.Config,
		logger: opts.Logger,
	}
}
