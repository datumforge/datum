// Package cmd is our cobra/viper cli implementation
package cmd

import (
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const appName = "datum"

var (
	cfgFile string
	logger  *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "A datum repo for graph apis",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config/."+appName+".yaml)")
	viperBindFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().Bool("pretty", false, "enable pretty (human readable) logging output")
	viperBindFlag("pretty", rootCmd.PersistentFlags().Lookup("pretty"))

	rootCmd.PersistentFlags().Bool("debug", false, "debug logging output")
	viperBindFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".datum" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".datum")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix("datum")
	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()

	logger = newLogger()

	if err == nil {
		logger.Infow("using config file", "file", viper.ConfigFileUsed())
	}
}

// viperBindFlag provides a wrapper around the viper bindings that panics if an error occurs
func viperBindFlag(name string, flag *pflag.Flag) {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}

// newLogger creates a new zap logger with the appropriate configuration based on the viper settings for pretty and debug
func newLogger() *zap.SugaredLogger {
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
