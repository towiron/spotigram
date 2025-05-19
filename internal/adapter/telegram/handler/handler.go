package handler

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/towiron/spotigram/internal/pkg/config"
	"go.uber.org/fx"
)

type Session struct {
	Scenario string
	Step     int
	Data     map[string]string
}

type Handler struct {
	config         config.Configer
	commands       map[string]func(*tgbotapi.BotAPI, *tgbotapi.Message) error
	defaultHandler func(*tgbotapi.BotAPI, *tgbotapi.Message) error

	sessions  map[int64]*Session
	scenarios map[string][]func(*Handler, *Session, *tgbotapi.BotAPI, *tgbotapi.Message) (bool, error)
	mu        sync.Mutex
}

type Options struct {
	fx.In
	Config config.Configer
}

var Module = fx.Provide(New)

func New(opts Options) *Handler {
	h := &Handler{
		config:    opts.Config,
		commands:  make(map[string]func(*tgbotapi.BotAPI, *tgbotapi.Message) error),
		sessions:  make(map[int64]*Session),
		scenarios: make(map[string][]func(*Handler, *Session, *tgbotapi.BotAPI, *tgbotapi.Message) (bool, error)),
	}

	return h
}
