package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var Logger *zap.Logger

func init() {
	// 错误日志
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{
		"logs/logfile.log",
	}

	var err error
	Logger, err = config.Build()
	if err != nil {
		log.Fatal("Failed to initialize error logger:", err)
	}
}
