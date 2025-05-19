package adapter

import (
	"github.com/towiron/spotigram/internal/adapter/telegram"
	"go.uber.org/fx"
)

var Module = fx.Options(
	telegram.Module,
)
