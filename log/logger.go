package log

type LoggerLevel int8

const (
	LevelDebug LoggerLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l LoggerLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	}
	return ""
}

type Logger interface {
	Log(Level LoggerLevel, kv ...interface{}) error
}
