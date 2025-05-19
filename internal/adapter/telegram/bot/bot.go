package bot

import (
	"context"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/towiron/spotigram/internal/adapter/telegram/handler"
	"github.com/towiron/spotigram/internal/pkg/config"
	"github.com/towiron/spotigram/internal/pkg/global"
	"github.com/towiron/spotigram/internal/pkg/logger"
	"go.uber.org/fx"
)

type Options struct {
	fx.In
	fx.Lifecycle
	Config  config.Configer
	Logger  logger.Logger
	Handler *handler.Handler
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(registerHooks),
)

func New(opts Options) (*tgbotapi.BotAPI, error) {
	token := opts.Config.String(global.TELEGRAM_BOT_TOKEN)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	opts.Logger.DebugF("Authorized on account %s", bot.Self.UserName)
	return bot, nil
}

func registerHooks(lc fx.Lifecycle, bot *tgbotapi.BotAPI, h *handler.Handler) {
	stopChan := make(chan struct{})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go startPolling(bot, h, stopChan)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			close(stopChan)

			time.Sleep(100 * time.Millisecond)
			log.Println("Telegram bot stopped")
			return nil
		},
	})
}

func startPolling(bot *tgbotapi.BotAPI, h *handler.Handler, stopChan chan struct{}) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for {
		select {
		case <-stopChan:
			return
		case upd := <-updates:
			if upd.Message == nil {
				continue
			}
		}
	}
}
