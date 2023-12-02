package logger

import (
	"fmt"
	"os"
)

func Info(a ...any) {
	fmt.Println(a...)
}

func Infof(msg string, a ...any) {
	Info(fmt.Sprintf(msg, a...))
}

func Fatal(msg string) {
	Infof(msg)
	os.Exit(1)
}

func FatalErr(msg string, err error) {
	Fatal(fmt.Sprintf("%v, err: %v", msg, err))
	PrintDebugTip()
}

func PrintDebugTip() {
	fmt.Println("For more information, run with --debug")
}
