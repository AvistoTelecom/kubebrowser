package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func InitLogger(logLevel string, isDev bool) error {
	zapLogger, err := newLogger(logLevel, isDev)
	if err != nil {
		return err
	}
	logger = zapLogger.Sugar()

	return nil
}

func newLogger(logLevel string, isDev bool) (*zap.Logger, error) {
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		fmt.Println("Unknown log level '%s', falling back to INFO", logLevel)
		level = zapcore.InfoLevel
	}

	if isDev {
		return zap.Config{
			Level:         zap.NewAtomicLevelAt(level),
			Encoding:      "console",
			OutputPaths:   []string{"stdout"},
			Development:   true,
			EncoderConfig: zap.NewDevelopmentEncoderConfig(),
		}.Build()
	}

	return zap.Config{
		Level:         zap.NewAtomicLevelAt(level),
		Encoding:      "json",
		OutputPaths:   []string{"stdout"},
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}.Build()
}
