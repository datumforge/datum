package logger

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

type Log interface {
	Info(message string)
	Warning(message string)
	Error(message string, err error)
	Fatal(message string, err error)
	Panic(message string, err error)
}

func NewLogger(conf Config) *Logger {
	rotator := &lumberjack.Logger{
		Filename:   conf.FileLocation,
		MaxSize:    conf.FileMaxSize, // megabytes
		MaxBackups: conf.FileMaxBackup,
		MaxAge:     conf.FileMaxAge, // days
	}

	encoderConfig := zap.NewDevelopmentEncoderConfig()

	if conf.Env == "production" {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "logLevel"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(
		jsonEncoder,
		zapcore.AddSync(rotator),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	if conf.Stdout {
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(
				consoleEncoder,
				zapcore.AddSync(os.Stdout),
				zap.NewAtomicLevelAt(zap.InfoLevel),
			),
		)
	}

	log := zap.New(core, zap.AddStacktrace(zap.ErrorLevel), zap.AddCallerSkip(2)).With( // nolint: gomnd
		zap.String("app", conf.App),
		zap.String("appVer", conf.AppVer),
		zap.String("env", conf.Env),
	)

	return &Logger{log}
}

func (l *Logger) Info(message string) {
	l.log.Info(message)
}

func (l *Logger) Warning(message string) {
	l.log.Warn(message)
}

func (l *Logger) Error(message string, err error) {
	l.log.Error(message, zap.Error(err))
}

func (l *Logger) Fatal(message string, err error) {
	l.log.Fatal(message, zap.Error(err))
}

func (l *Logger) Panic(message string, err error) {
	l.log.Panic(message, zap.Error(err))
}
