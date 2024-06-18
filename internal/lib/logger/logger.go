package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	envLocal = "local"
	envProd  = "production"
)

func New(env string) *zap.Logger {
	var log *zap.Logger
	var core zapcore.Core

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleWriter := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, zap.DebugLevel)

	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	file, err := os.Create("logs.log")
	if err != nil {
		panic(err)
	}
	fileWriter := zapcore.AddSync(file)
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, zap.InfoLevel)

	switch env {
	case envLocal:
		core = zapcore.NewTee(consoleCore)
	case envProd:
		core = zapcore.NewTee(consoleCore, fileCore)
	default:
		core = zapcore.NewTee(consoleCore, fileCore)
	}

	log = zap.New(core)

	return log
}
