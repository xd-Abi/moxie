package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	internalLogger *zap.Logger
}

func New() *Log {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeCaller = nil
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()

	return &Log{internalLogger: logger}
}

func (l *Log) Fatal(msg string, args ...interface{}) {
	l.internalLogger.Sugar().Fatalf(msg, args...)
}

func (l *Log) Error(msg string, args ...interface{}) {
	l.internalLogger.Sugar().Errorf(msg, args...)
}

func (l *Log) Warn(msg string, args ...interface{}) {
	l.internalLogger.Sugar().Warnf(msg, args...)
}

func (l *Log) Info(msg string, args ...interface{}) {
	l.internalLogger.Sugar().Infof(msg, args...)
}
