package log

// 栈跳过的默认层级
const defaultSkip = 3

// 配置文件的日志级别
const (
	levelCfgDebug = "debug"
	levelCfgInfo  = "info"
	levelCfgWarn  = "warn"
	levelCfgError = "error"
	levelCfgFatal = "fatal"
)

// 日志级别名称
const (
	levelNameDebug   = "DEBUG"
	levelNameInfo    = "INFO"
	levelNameWarn    = "WARN"
	levelNameError   = "ERROR"
	levelNameFatal   = "FATAL"
	levelNameUnknown = "UNKNOWN"
)
