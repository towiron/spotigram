package telegram

import (
	"github.com/towiron/spotigram/internal/adapter/telegram/bot"
	"github.com/towiron/spotigram/internal/adapter/telegram/handler"
	"go.uber.org/fx"
)

var Module = fx.Options(
	bot.Module,
	handler.Module,
)
