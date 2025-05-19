package pkg

import (
	"github.com/towiron/spotigram/internal/pkg/config"
	"github.com/towiron/spotigram/internal/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
)
