package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.Logger
}

func (z *zapLogger) Debug(msg string, args ...LogStruct) {
	z.logger.Debug(msg, z.parseArgs(args)...)
}

func (z *zapLogger) Info(msg string, args ...LogStruct) {
	z.logger.Info(msg, z.parseArgs(args)...)

}

func (z *zapLogger) Warn(msg string, args ...LogStruct) {
	z.logger.Warn(msg, z.parseArgs(args)...)
}

func (z *zapLogger) Error(msg string, args ...LogStruct) {
	z.logger.Error(msg, z.parseArgs(args)...)
}

func (z *zapLogger) parseArgs(args []LogStruct) []zapcore.Field {
	result := make([]zapcore.Field, len(args))
	for i, arg := range args {
		result[i] = zap.String(arg.key, fmt.Sprintf("%v", arg.val))
	}
	return result
}

func newLogger(lvl LogLevel) Logger {

	level := zap.DebugLevel
	switch lvl {
	case INFO:
		level = zap.InfoLevel
	case WARN:
		level = zap.WarnLevel
	case ERROR:
		level = zap.WarnLevel
	default:
	}

	disableStackTrace := true
	if lvl == DEBUG {
		disableStackTrace = false
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: disableStackTrace,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return &zapLogger{
		logger: zap.Must(config.Build()),
	}
}
