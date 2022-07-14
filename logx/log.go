package logx

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type FileOption struct {
	// https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2?utm_source=godoc#Logger
	Filename string
	MaxSize  int
	MaxAge   int
	Compress bool

	Duration time.Duration // set 0 means don't rotate by duration
}

type LoggerX struct {
	*zap.SugaredLogger
}

// Init init logx
// if file is nil only output log info to os.Stdout
func Init(file *FileOption) {
	// console encoder and writer
	consoleEncoder, consoleWriter := getConsole()

	// json encoder and writer
	var (
		jsonEncoder zapcore.Encoder
		jsonWriter  zapcore.WriteSyncer
	)
	if file != nil {
		jsonEncoder, jsonWriter = getJSON(file)
	}

	// level
	atom := zap.NewAtomicLevel()

	// core
	cores := []zapcore.Core{
		zapcore.NewCore(consoleEncoder, consoleWriter, atom),
	}
	if file != nil {
		cores = append(cores, zapcore.NewCore(jsonEncoder, jsonWriter, atom))
	}
	core := zapcore.NewTee(cores...)

	// logger
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	).Sugar()

	res := &LoggerX{
		SugaredLogger: logger,
	}

	// set package global variable
	level = atom
	loggerx = res
}

func getConsole() (zapcore.Encoder, zapcore.WriteSyncer) {
	encoder := zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	)
	writer := zapcore.AddSync(os.Stdout)

	return encoder, writer
}

func getJSON(file *FileOption) (zapcore.Encoder, zapcore.WriteSyncer) {
	encoder := zapcore.NewJSONEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	)

	// https://github.com/uber-go/zap/blob/master/FAQ.md#does-zap-support-log-rotation
	r := &lumberjack.Logger{
		Filename:  file.Filename,
		MaxSize:   file.MaxSize,
		MaxAge:    file.MaxAge,
		Compress:  file.Compress,
		LocalTime: true,
	}

	// https://github.com/natefinch/lumberjack/issues/17#issuecomment-185846531
	if file.Duration != 0 {
		go func() {
			for {
				<-time.After(file.Duration)
				if err := r.Rotate(); err != nil {
					log.Printf("[ERROR] [logx] rotate err: %s", err.Error())
				}
			}
		}()
	}
	writer := zapcore.AddSync(r)

	return encoder, writer
}
