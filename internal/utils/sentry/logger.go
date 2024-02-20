package sentry

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new zap logger with the appropriate configuration based on the viper settings for pretty and debug
func NewLogger() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	if viper.GetBool("pretty") {
		cfg = zap.NewDevelopmentConfig()
	}

	if viper.GetBool("debug") {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = logger.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, &SentryZapCore{enabledLevel: zapcore.DebugLevel})
	}))

	return logger.Sugar()
}

// zapLevelToSentryLevel converts a zap log level to a sentry log level
func zapLevelToSentryLevel(level zapcore.Level) sentry.Level {
	switch level {
	case zapcore.DebugLevel:
		return sentry.LevelDebug
	case zapcore.InfoLevel:
		return sentry.LevelInfo
	case zapcore.WarnLevel:
		return sentry.LevelWarning
	case zapcore.ErrorLevel:
		return sentry.LevelError
	case zapcore.DPanicLevel:
		return sentry.LevelFatal
	case zapcore.PanicLevel:
		return sentry.LevelFatal
	case zapcore.FatalLevel:
		return sentry.LevelFatal
	default:
		return sentry.LevelFatal
	}
}

// SentryZapCore is a zap core that sends logs to sentry
type SentryZapCore struct {
	enabledLevel zapcore.Level
	fields       map[string]interface{}
}

// Enabled returns true if the log level is enabled
func (s *SentryZapCore) Enabled(level zapcore.Level) bool {
	return s.enabledLevel <= level
}

// With returns a new core with the fields added
func (s *SentryZapCore) With(fields []zapcore.Field) zapcore.Core {
	copied := make(map[string]interface{}, len(s.fields))
	for k, v := range s.fields {
		copied[k] = v
	}

	encoder := zapcore.NewMapObjectEncoder()
	for _, f := range fields {
		f.AddTo(encoder)
	}

	for k, v := range encoder.Fields {
		copied[k] = v
	}

	return &SentryZapCore{fields: copied, enabledLevel: s.enabledLevel}
}

// Check returns a checked entry if the log level is enabled for the core
func (s *SentryZapCore) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if s.Enabled(entry.Level) {
		checkedEntry.AddCore(entry, s)
	}

	return checkedEntry
}

// Write sends the log entry to sentry
func (s *SentryZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	event := sentry.NewEvent()
	event.Message = entry.Message
	event.Timestamp = time.Now()
	event.Level = zapLevelToSentryLevel(entry.Level)
	event.Platform = "Golang"
	exceptions := make([]sentry.Exception, 0)

	for _, f := range fields {
		if f.Type == zapcore.ErrorType {
			err := f.Interface.(error)

			trace := sentry.ExtractStacktrace(err)

			if trace == nil {
				trace = sentry.NewStacktrace()
			}

			exceptions = append(exceptions, sentry.Exception{
				Type:       entry.Message,
				Value:      entry.Caller.TrimmedPath(),
				Stacktrace: trace,
			})
		}
	}

	event.Exception = exceptions

	sentry.CaptureEvent(event)

	err := s.Sync()

	if err != nil {
		return err
	}

	return nil
}

// Sync flushes the sentry event
func (s *SentryZapCore) Sync() error {
	sentry.Flush(2 * time.Second) // nolint:gomnd
	return nil
}
