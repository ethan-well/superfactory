package logger

import (
	"fmt"
	"log"
	"runtime"
)

var (
	Info = LogBasic("info", Teal)
	Waring = LogBasic("warn", Yellow)
	Error = LogBasic("error", Red)

	Infof = LogBasicf("info", Teal)
	Waringf = LogBasicf("warn", Yellow)
	Errorf = LogBasicf("error", Red)
)

func LogBasic(level string, colorFunc func(string2 string) string) func(...interface{}) {
	return func(args ...interface{}) {
		pc, fn, line, _ := runtime.Caller(1)

		str := fmt.Sprintf("[%s] in %s [%s:%d] %s\n", level, runtime.FuncForPC(pc).Name(), fn, line, fmt.Sprint(args))

		log.Printf(colorFunc(str))
	}
}

func LogBasicf(level string, colorFunc func(string2 string) string) func(string, ...interface{}) {
	return func(format string, args ...interface{}) {
		pc, fn, line, _ := runtime.Caller(1)

		str := fmt.Sprintf("[%s] in %s [%s:%d] %s\n", level, runtime.FuncForPC(pc).Name(), fn, line, fmt.Sprintf(format, args))

		log.Printf(colorFunc(str))
	}
}