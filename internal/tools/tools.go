package tools

import (
	"runtime"
)

func CurrentFunction() string {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		return ""
	}

	return runtime.FuncForPC(counter).Name()
}
