package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerx *LoggerX
	level   zap.AtomicLevel
)

type levelType = zapcore.Level

var (
	DebugLevel = zap.DebugLevel
	InfoLevel  = zap.InfoLevel
	WarnLevel  = zap.WarnLevel
	ErrorLevel = zap.ErrorLevel
	PanicLevel = zapcore.PanicLevel
)

func SetLevel(l levelType) {
	level.SetLevel(l)
}

func Sync() error {
	return loggerx.Sync()
}

func Debug(args ...any) {
	loggerx.Debug(args...)
}

func Debugf(template string, args ...any) {
	loggerx.Debugf(template, args...)
}

func Info(args ...any) {
	loggerx.Info(args...)
}

func Infof(template string, args ...any) {
	loggerx.Infof(template, args...)
}

func Warn(args ...any) {
	loggerx.Warn(args...)
}

func Warnf(template string, args ...any) {
	loggerx.Warnf(template, args...)
}

func Error(args ...any) {
	loggerx.Error(args...)
}

func Errorf(template string, args ...any) {
	loggerx.Errorf(template, args...)
}

func Panic(args ...any) {
	loggerx.Panic(args...)
}

func Panicf(template string, args ...any) {
	loggerx.Panicf(template, args...)
}
