package middleware

import (
	"cuiliang/utils/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggerMiddleware(c *gin.Context) {
	start := time.Now()

	// Process request
	c.Next()

	end := time.Now()
	latency := end.Sub(start)

	clientIP := c.ClientIP()
	method := c.Request.Method
	path := c.Request.URL.Path
	statusCode := c.Writer.Status()
	userAgent := c.Request.UserAgent()

	// Log fields
	logFields := []zap.Field{
		zap.String("clientIP", clientIP),
		zap.String("method", method),
		zap.String("path", path),
		zap.String("protocol", c.Request.Proto),
		zap.Int("statusCode", statusCode),
		zap.Int("responseSize", c.Writer.Size()),
		zap.String("referer", c.Request.Referer()),
		zap.String("userAgent", userAgent),
		zap.Duration("latency", latency),
	}

	// Use Zap logger to log the message
	logger.AccessLogger.Info("Request handled", logFields...)

	// 记录错误日志
	if len(c.Errors) > 0 {
		for _, err := range c.Errors.Errors() {
			logger.ErrorLogger.Error("Error Log", zap.Error(errors.New(err)))
		}
	}
}
