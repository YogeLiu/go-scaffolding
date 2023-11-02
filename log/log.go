package log

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Debug 日志
func Debug(c *gin.Context, a ...interface{}) {
	if gin.Mode() == gin.ReleaseMode {
		return
	}

	if logger == nil {
		return
	}

	content := logger.format(c, fmt.Sprint(a...), defaultSkip)
	logger.Debug(content)
}

// Info 日志
func Info(c *gin.Context, a ...interface{}) {
	if logger == nil {
		return
	}

	content := logger.format(c, fmt.Sprint(a...), defaultSkip)
	logger.Info(content)
}

// Warn 日志
func Warn(c *gin.Context, a ...interface{}) {
	if logger == nil {
		return
	}

	content := logger.format(c, fmt.Sprint(a...), defaultSkip)
	logger.Warn(content)
}

// Error 日志
func Error(c *gin.Context, a ...interface{}) {
	if logger == nil {
		return
	}

	content := logger.format(c, fmt.Sprint(a...), defaultSkip)
	logger.Error(content)
}

// Fatal 日志
func Fatal(c *gin.Context, a ...interface{}) {
	if logger == nil {
		return
	}

	content := logger.format(c, fmt.Sprint(a...), defaultSkip)
	logger.Fatal(content)
}
