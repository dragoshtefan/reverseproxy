package log

// log levels in order , from most verbose to least :
//Trace
//Debug
//Info
//Warn
//Error
//Fatal
// we expose only debug, info , warn and error for ease of use

type LogStruct struct {
	key string
	val interface{}
}
type Logger interface {
	Debug(msg string, args ...LogStruct)
	Info(msg string, args ...LogStruct)
	Warn(msg string, args ...LogStruct)
	Error(msg string, args ...LogStruct)
}

func NewLogger(logLevel LogLevel) Logger {
	return newLogger(logLevel)
}

type LogLevel int

const (
	DEBUG LogLevel = -1
	INFO  LogLevel = 0
	WARN  LogLevel = 1
	ERROR LogLevel = 2
)
