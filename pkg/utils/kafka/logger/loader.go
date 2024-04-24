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
		singleton = NewLogger(conf)
	})

	return singleton
}

func Info(message string) {
	singleton.Info(message)
}

func Warning(message string) {
	singleton.Warning(message)
}

func Error(message string, err error) {
	singleton.Error(message, err)
}

func Fatal(message string, err error) {
	singleton.Fatal(message, err)
}

func Panic(message string, err error) {
	singleton.Panic(message, err)
}
