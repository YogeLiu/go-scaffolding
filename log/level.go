package log

import (
	"errors"

	"go.uber.org/zap/zapcore"
)

// checkLevel 检测日志级别
func checkLevel(level string) error {
	switch level {
	case levelCfgDebug:
		return nil
	case levelCfgInfo:
		return nil
	case levelCfgWarn:
		return nil
	case levelCfgError:
		return nil
	case levelCfgFatal:
		return nil
	}

	return errors.New("log level error")
}

// getLevel 获取日志级别
func getLevel(level string) zapcore.Level {
	switch level {
	case levelCfgDebug:
		return zapcore.DebugLevel
	case levelCfgInfo:
		return zapcore.InfoLevel
	case levelCfgWarn:
		return zapcore.WarnLevel
	case levelCfgError:
		return zapcore.ErrorLevel
	case levelCfgFatal:
		return zapcore.FatalLevel
	}

	return zapcore.InfoLevel
}
