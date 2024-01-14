package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/jigadhirasu/rex"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger level能設置的等級
// 空值會設置為info
// var validLevel = []string{
// 	"debug", "DEBUG", // -1
// 	"info", "INFO", "", // 0
// 	"warn", "WARN", // 1
// 	"error", "ERROR", // 2
// 	"dpanic", "DPANIC", // 3
// 	"panic", "PANIC", // 4
// 	"fatal", "FATAL", // 5
// }

func getPath() string {
	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		logPath = "/var/log"
	}

	projectName := os.Getenv("PROJECT_NAME")
	if projectName == "" {
		projectName = "default"
	}

	path := fmt.Sprintf("%s/%s", logPath, projectName)
	os.MkdirAll(path, 0644)

	return path
}

func Out(ctxx ...rex.Context) *zap.Logger {
	if len(ctxx) < 1 {
		ctxx = append(ctxx, rex.NewContext(context.TODO()))
	}
	ctx := ctxx[0]
	if log, ok := rex.Maybe[*zap.Logger](ctx, CtxOutLogger); ok {
		return log
	}

	path := getPath()

	outfile := path + "/out.log"
	errfile := path + "/err.log"

	// logtext := env.Getenv("LOG_LEVEL")
	// lv, err := zapcore.ParseLevel(logtext)
	// if err != nil {
	// 	panic(fmt.Sprintf(`env LOG_LEVEL [%s] invalid, need be %v`, logtext, validLevel))
	// }

	core := zapcore.NewTee(
		zap.Must(ConsoleConfig(zap.DebugLevel).Build()).Core(),
		zap.Must(FileConfig(zap.InfoLevel, outfile).Build()).Core(),
		// zap.Must(ConsoleConfig(zap.ErrorLevel)(development).Build()).Core(),
		zap.Must(FileConfig(zap.ErrorLevel, errfile).Build()).Core(),
	)

	// 開發環境
	if os.Getenv("ENV") == "develop" {
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel), zap.Development())
	}

	// 非開發環境
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

func Access(ctxx ...rex.Context) *zap.Logger {
	if len(ctxx) < 1 {
		ctxx = append(ctxx, rex.NewContext(context.TODO()))
	}
	ctx := ctxx[0]
	if log, ok := rex.Maybe[*zap.Logger](ctx, CtxAccessLogger); ok {
		return log
	}

	path := getPath()

	accessfile := path + "/access.log"

	core := zapcore.NewTee(
		zap.Must(ConsoleConfig(zap.InfoLevel).Build()).Core(),
		zap.Must(FileConfig(zap.InfoLevel, accessfile).Build()).Core(),
	)

	return zap.New(core, zap.AddCaller())

}
