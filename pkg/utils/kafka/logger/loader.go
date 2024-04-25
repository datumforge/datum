package logger

import (
	"sync"
)

var (
	once      sync.Once
	singleton Log
)

func Load(conf Config) Log {
	once.Do(func() {
		singleton = NewZapLogger(conf)
	})

	return singleton
}

func Info(message string) {
	singleton.Info(message, nil)
}

func Warning(message string) {
	singleton.Debug(message, nil)
}

func Error(message string, err error) {
	singleton.Error(message, err, nil)
}

func Fatal(message string, err error) {
	singleton.Trace(message, nil)
}
