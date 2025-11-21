package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZap(opts ...Option) *ZapLogger {
	options := &options{
		output:   os.Stdout,
		logLevel: zapcore.WarnLevel,
	}
	for _, apply := range opts {
		apply(options)
	}

	encoder := zapcore.NewJSONEncoder(newEncoderConfig())
	sink := zapcore.Lock(zapcore.AddSync(options.output))
	level := zap.NewAtomicLevelAt(options.logLevel)
	core := zapcore.NewCore(encoder, sink, level)

	logger := zap.New(
		core,
		zap.ErrorOutput(sink),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	).Sugar()

	if options.env != nil {
		logger = logger.With(EnvKey, *options.env)
	}

	if options.pid != nil {
		logger = logger.With(PidKey, *options.pid)
	}

	if options.name != nil {
		logger = logger.With(NameKey, *options.name)
	}

	if options.hostname != nil {
		logger = logger.With(InstKey, *options.hostname)
	}

	return &ZapLogger{
		logger: logger,
	}
}

func encodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(l.String())
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey: MsgKey,

		LevelKey:    LevelKey,
		EncodeLevel: encodeLevel,

		TimeKey:    TimeKey,
		EncodeTime: zapcore.ISO8601TimeEncoder,
	}
}

func (l *ZapLogger) With(key string, value any) Logger {
	return &ZapLogger{
		logger: l.logger.With(key, value),
	}
}

func (l *ZapLogger) WithError(err error) Logger {
	return l.With("error", err)
}

func (l *ZapLogger) Debugf(msg string, args ...any) {
	l.logger.Debugf(msg, args...)
}

func (l *ZapLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *ZapLogger) Infof(msg string, args ...any) {
	l.logger.Infof(msg, args...)
}

func (l *ZapLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *ZapLogger) Warnf(msg string, args ...any) {
	l.logger.Warnf(msg, args...)
}

func (l *ZapLogger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *ZapLogger) Errorf(msg string, args ...any) {
	l.logger.Errorf(msg, args...)
}

func (l *ZapLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *ZapLogger) Fatalf(msg string, args ...any) {
	l.logger.Fatalf(msg, args...)
}

func (l *ZapLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}
