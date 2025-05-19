package logger

import (
	"time"

	"github.com/towiron/spotigram/internal/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	_CFG_MODE  = "mode"
	_PROD_MODE = "prod"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)

	DebugF(format string, args ...any)
	InfoF(format string, args ...any)
	WarnF(format string, args ...any)
	ErrorF(format string, args ...any)
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
	var (
		baseLogger *zap.Logger
		err        error
	)

	switch opts.Config.String(_CFG_MODE) {
	case _PROD_MODE:
		baseLogger, err = newProdConfig().Build()
	default:
		baseLogger, err = newDevConfig().Build()
	}

	if err != nil {
		return nil, err
	}

	return &zapLogger{Logger: baseLogger}, nil
}

func newProdConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "level",
			TimeKey:    "ts",
			CallerKey:  "caller",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.RFC3339))
			},
			EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(level.String())
			},
			EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(caller.TrimmedPath())
			},
		},
	}
}

func newDevConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "level",
			TimeKey:    "ts",
			CallerKey:  "caller",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.RFC3339))
			},
			EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(level.String())
			},
			EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(caller.TrimmedPath())
			},
		},
	}
}

func (l *zapLogger) Debug(msg string, fields ...zap.Field) { l.Logger.Debug(msg, fields...) }
func (l *zapLogger) Info(msg string, fields ...zap.Field)  { l.Logger.Info(msg, fields...) }
func (l *zapLogger) Warn(msg string, fields ...zap.Field)  { l.Logger.Warn(msg, fields...) }
func (l *zapLogger) Error(msg string, fields ...zap.Field) { l.Logger.Error(msg, fields...) }

func (l *zapLogger) DebugF(format string, args ...any) {
	l.Logger.Sugar().Debugf(format, args...)
}

func (l *zapLogger) InfoF(format string, args ...any) {
	l.Logger.Sugar().Infof(format, args...)
}

func (l *zapLogger) WarnF(format string, args ...any) {
	l.Logger.Sugar().Warnf(format, args...)
}

func (l *zapLogger) ErrorF(format string, args ...any) {
	l.Logger.Sugar().Errorf(format, args...)
}
