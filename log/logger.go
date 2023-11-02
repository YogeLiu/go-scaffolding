package log

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"scaffolding/conf"
	"scaffolding/middleware"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志
type Logger struct {
	*zap.Logger
}

var (
	logger *Logger
)

func Init() {
	if logger != nil {
		return
	}
	err := newLogger(&conf.Conf.Log)
	if err != nil {
		panic(err)
	}
}

// GetLogger 获取日志对象
func GetLogger() *Logger {
	return logger
}

// newLogger 初始化日志
func newLogger(cfg *conf.LogConfig) error {
	err := checkLevel(cfg.Level)
	if err != nil {
		return err
	}

	// 创建日志目录文件
	fileDir, fileName := filepath.Split(cfg.FilePath)
	if fileDir == "" || fileName == "" {
		return errors.New("log file path error")
	}

	err = createFileDir(fileDir)
	if err != nil {
		return err
	}

	// 按天分割
	fileExt := filepath.Ext(fileName)
	filePath := fileDir + "/" + strings.TrimSuffix(fileName, fileExt) + "_%Y%m%d.log"
	hook, err := rotatelogs.New(filePath)
	if err != nil {
		return err
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.LineEnding = zapcore.DefaultLineEnding
	encoderCfg.EncodeLevel = levelEncoder
	encoderCfg.EncodeTime = timeEncoder
	encoderCfg.EncodeDuration = durationEncoder
	encoderCfg.EncodeCaller = nil
	encoderCfg.EncodeName = zapcore.FullNameEncoder

	z := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.AddSync(hook),
			getLevel(cfg.Level),
		),
		//zap.AddCaller(),
		//zap.AddCallerSkip(1),
	)
	defer func() {
		_ = z.Sync()
	}()

	logger = &Logger{
		z,
	}

	return nil
}

// createFileDir 创建日志目录
func createFileDir(fileDir string) error {
	_, err := os.Stat(fileDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// getLevelName 获取日志级别
func getLevelName(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return levelNameDebug
	case zapcore.InfoLevel:
		return levelNameInfo
	case zapcore.WarnLevel:
		return levelNameWarn
	case zapcore.ErrorLevel:
		return levelNameError
	case zapcore.DPanicLevel:
		return levelNameError
	case zapcore.PanicLevel:
		return levelNameFatal
	case zapcore.FatalLevel:
		return levelNameFatal
	}

	return levelNameUnknown
}

// levelEncoder 日志级别格式
func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + getLevelName(l) + "]")
}

// timeEncoder 时间格式
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + time.Now().Format("2006-01-02 15:04:05.999") + "]")
}

// durationEncoder 时间格式
func durationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + d.String() + "]")
}

// format 格式化日志
// [time] [level] [ip] [uri] [duration] [trace_id] [file:line] [msg]
func (log *Logger) format(c *gin.Context, msg string, skip int) string {
	var bt bytes.Buffer
	bt.WriteString("[")
	if c != nil {
		bt.WriteString(c.Request.RemoteAddr)
	} else {
		bt.WriteString("")
	}

	bt.WriteString("] [")
	if c != nil {
		bt.WriteString(c.Request.RequestURI)
	} else {
		bt.WriteString("")
	}

	bt.WriteString("] [")
	if c != nil {
		bt.WriteString(log.getDuration(c))
	} else {
		bt.WriteString("")
	}

	bt.WriteString("] [")
	if c != nil {
		bt.WriteString(middleware.GetTraceId(c))
	} else {
		bt.WriteString("")
	}

	bt.WriteString("] [")
	bt.WriteString(log.getFuncFileLine(skip))
	bt.WriteString("] [")
	bt.WriteString(msg)
	bt.WriteString("]")

	return bt.String()
}

// getDuration 持续时间
func (log *Logger) getDuration(c *gin.Context) string {
	duration := float64(time.Since(middleware.GetRequestStartTime(c))) / float64(time.Second)
	return strconv.FormatFloat(duration, 'f', 3, 64)
}

// getFuncFileLine 获取代码方法文件行数
func (log *Logger) getFuncFileLine(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	var bt bytes.Buffer
	bt.WriteString(file)
	bt.WriteString(":")
	bt.WriteString(strconv.Itoa(line))

	return bt.String()
}
