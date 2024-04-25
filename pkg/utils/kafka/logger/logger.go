package logger

import (
	"os"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wrapper around the zap logger
type Logger struct {
	log    *zap.Logger
	fields watermill.LogFields
}

// Log is the interface for the logger
type Log interface {
	Info(message string, fields watermill.LogFields)
	Error(message string, err error, fields watermill.LogFields)
	Debug(message string, fields watermill.LogFields)
	Trace(message string, fields watermill.LogFields)
	With(fields watermill.LogFields) watermill.LoggerAdapter
}

func NewLogger(z *Logger) watermill.LoggerAdapter {
	return &Logger{log: z.log}
}

// NewLogger is a function that creates and configures a new logger instance using the zap
// logging library. It takes a `Config` struct as a parameter, which contains configuration settings
// for the logger such as file location, maximum file size, maximum backups, maximum age, environment,
// and whether to log to stdout
func NewZapLogger(conf Config) *Logger {
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

	return &Logger{log: log}
}

// Info writes info log with message and some fields.
func (l *Logger) Info(msg string, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.log.Info(msg, fs...)
}

// Error writes error log with message, error and some fields.
func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	fs = append(fs, zap.Error(err))
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.log.Error(msg, fs...)
}

// Debug writes debug log with message and some fields
func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.log.Debug(msg, fs...)
}

// Trace writes debug log instead of trace log because zap does not support trace level logging
func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.log.Debug(msg, fs...)
}

// With returns new LoggerAdapter with passed fields.
func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{
		log:    l.log,
		fields: l.fields.Add(fields),
	}
}
