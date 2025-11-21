package logger

import "go.uber.org/zap/zapcore"

type Logger interface {
	With(key string, value any) Logger
	WithError(err error) Logger

	Debugf(msg string, args ...any)
	Debug(msg string)

	Infof(msg string, args ...any)
	Info(msg string)

	Warnf(msg string, args ...any)
	Warn(msg string)

	Errorf(msg string, args ...any)
	Error(msg string)

	Fatalf(msg string, args ...any)
	Fatal(msg string)
}

func New(config *Config) Logger {
	opts := []Option{
		WithEnv(config.Env),
		WithName(config.Name),
		WithHostname(),
	}

	if config.Env == "local" {
		opts = append(opts, WithPID(), WithLogLevel(zapcore.DebugLevel))
	} else {
		opts = append(opts, WithLogLevel(zapcore.InfoLevel))
	}

	return NewZap(opts...)
}
