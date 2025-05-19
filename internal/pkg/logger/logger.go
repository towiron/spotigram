package logger

import (
	"github.com/towiron/spotigram/internal/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	_CFG_MODE  = "mode"
	_PROD_MODE = "prod"
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
}

type zapLogger struct {
	*zap.Logger
}

type Options struct {
	fx.In
	Config config.Configer
}

var Module = fx.Provide(New)

func New(opts Options) (Logger, error) {
	if opts.Config.String(_CFG_MODE) == _PROD_MODE {
		log, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		return &zapLogger{log}, nil
	}

	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &zapLogger{log}, nil
}

func (l *zapLogger) Info(msg string, fields ...zap.Field)  { l.Logger.Info(msg, fields...) }
func (l *zapLogger) Error(msg string, fields ...zap.Field) { l.Logger.Error(msg, fields...) }
func (l *zapLogger) Debug(msg string, fields ...zap.Field) { l.Logger.Debug(msg, fields...) }
func (l *zapLogger) Warn(msg string, fields ...zap.Field)  { l.Logger.Warn(msg, fields...) }
