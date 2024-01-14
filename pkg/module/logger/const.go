package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogFileFlag             = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	LogFilePerm os.FileMode = 0644
)

const (
	// 輸出日誌的物件
	CtxOutLogger    = "ctx_out_logger"
	CtxAccessLogger = "ctx_access_logger"
	CtxKeyRequestID = "ctx_request_id"
	CtxUUID         = "ctx_uuid"
	CtxIP           = "ctx_ip"
	CtxVersion      = "ctx_version"
	CtxStartTime    = "ctx_start_at"
)

type FileName string

const (
	OutLog    FileName = "out.log"
	ErrLog    FileName = "err.log"
	AccessLog FileName = "access.log"
	DebugLog  FileName = "debug.log"
)

func ConsoleConfig(lv zapcore.Level) zap.Config {

	return zap.Config{
		Level:            zap.NewAtomicLevelAt(lv),
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},

		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}

func FileConfig(lv zapcore.Level, filepath ...string) zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(lv),
		Encoding:         "json",
		OutputPaths:      filepath,
		ErrorOutputPaths: filepath,

		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}
