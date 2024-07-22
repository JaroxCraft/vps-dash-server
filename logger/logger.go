package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var logger *zap.SugaredLogger

func Get() *zap.SugaredLogger {
	if logger == nil {
		create()
	}
	return logger
}

func create() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.DebugLevel)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.Kitchen)

	created, err := config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %s", err)
	}

	logger = created.Sugar()
	return logger
}
