package log

import (
	"fmt"

	"go.uber.org/zap"
)

type ILogger interface {
	Flush()
	WithField(key string, value interface{}) ILogger
	WithPrefix(prefix string) ILogger

	Info(args ...interface{})
	Infof(template string, args ...interface{})

	Debug(args ...interface{})
	Debugf(template string, args ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type Logger struct {
	internal *zap.Logger
}

func (l *Logger) Flush() {
	err := l.internal.Sync()
	if err != nil {
		_ = fmt.Errorf("log could not flush, error: %+v", err)
	}
}

func (l *Logger) WithField(key string, value interface{}) ILogger {
	return &Logger{
		internal: l.internal.With(
			zap.Any(key, value),
		),
	}
}

func (l *Logger) WithPrefix(prefix string) ILogger {
	return l.WithField("prefix", prefix)
}

func (l *Logger) WithRequestID(requestID string) ILogger {
	return l.WithField("request_id", requestID)
}

func (l *Logger) Info(args ...interface{}) {
	l.internal.Sugar().Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.internal.Sugar().Infof(template, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.internal.Sugar().Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.internal.Sugar().Debugf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.internal.Sugar().Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.internal.Sugar().Errorf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.internal.Sugar().Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.internal.Sugar().Fatalf(template, args...)
}

func NewLogger() ILogger {
	internal, _ := zap.NewProduction()

	logger := &Logger{
		internal: internal,
	}

	return logger
}

func ProvideLogger() (ILogger, func(), error) {
	logger := NewLogger()

	cleanup := func() {
		logger.Flush()
	}

	return logger, cleanup, nil
}
