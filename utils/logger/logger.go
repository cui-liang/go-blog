package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var AccessLogger *zap.Logger
var ErrorLogger *zap.Logger

func init() {
	// 访问日志
	accessCfg := zap.NewDevelopmentConfig()
	accessCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	accessCfg.OutputPaths = []string{
		"./logs/access_" + time.Now().Format("2006-01-02") + ".log",
	}

	var err error
	AccessLogger, err = accessCfg.Build()
	if err != nil {
		log.Fatal("Failed to initialize access logger:", err)
	}

	// 错误日志
	errorCfg := zap.NewDevelopmentConfig()
	errorCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	errorCfg.OutputPaths = []string{
		"./logs/error_" + time.Now().Format("2006-01-02") + ".log",
	}

	ErrorLogger, err = errorCfg.Build()
	if err != nil {
		log.Fatal("Failed to initialize error logger:", err)
	}
}
