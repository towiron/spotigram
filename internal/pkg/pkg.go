package pkg

import (
	"github.com/towiron/spotigram/internal/pkg/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
)
