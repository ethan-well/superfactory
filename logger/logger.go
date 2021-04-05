package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
)

var (
	Info   = LogBasic("info", Teal)
	Waring = LogBasic("warn", Yellow)
	Error  = LogBasic("error", Red)

	Infof   = LogBasicf("info", Teal)
	Waringf = LogBasicf("warn", Yellow)
	Errorf  = LogBasicf("error", Red)
)

func LogBasic(level string, colorFunc func(string2 string) string) func(...interface{}) {
	return func(args ...interface{}) {
		pc, fn, line, ok := runtime.Caller(1)
		if !ok {
			fn = "??? file missing ???"
		}

		str := fmt.Sprintf("[%s] in %s [%s:%d] %s\n", level, runtime.FuncForPC(pc).Name(), fn, line, fmt.Sprint(args))

		log.Printf(colorFunc(str))
	}
}

func LogBasicf(level string, colorFunc func(string2 string) string) func(string, ...interface{}) {
	return func(format string, args ...interface{}) {
		pc, fn, line, ok := runtime.Caller(1)
		if !ok {
			fn = "??? file missing ???"
		}

		str := fmt.Sprintf("[%s] in %s [%s:%d] %s\n", level, runtime.FuncForPC(pc).Name(), fn, line, fmt.Sprintf(format, args))

		log.Printf(colorFunc(str))
	}
}

func ToJson(log interface{}) string {
	by, err := json.Marshal(log)
	if err != nil {
		return fmt.Sprintf("%v", log)
	}

	return string(by)
}
