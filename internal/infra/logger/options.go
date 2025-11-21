package logger

import (
	"io"
	"os"

	"github.com/samber/lo"
	"go.uber.org/zap/zapcore"
)

type options struct {
	env      *string
	pid      *int
	name     *string
	hostname *string
	output   io.Writer
	logLevel zapcore.Level
}

type Option func(*options)

func WithEnv(env string) Option {
	return func(o *options) {
		o.env = &env
	}
}

func WithPID() Option {
	return func(o *options) {
		o.pid = lo.ToPtr(os.Getpid())
	}
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = &name
	}
}

func WithHostname() Option {
	return func(o *options) {
		hostname, err := os.Hostname()
		if err != nil {
			return
		}
		o.hostname = &hostname
	}
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func WithLogLevel(logLevel zapcore.Level) Option {
	return func(o *options) {
		o.logLevel = logLevel
	}
}
