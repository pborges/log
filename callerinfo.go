package log

import (
	"runtime"
	"strings"
	"strconv"
)

type callerInfo struct {
	File    string
	Line    int
	Package string
	Func    string
}

func getCallerInfo(callSkip int) (callerInfo) {
	pc, file, line, _ := runtime.Caller(callSkip + 1) // plus one is to account for itself

	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)

	funcStartOffset := pl - 1
	if _, err := strconv.Atoi(parts[funcStartOffset]); err == nil && pl >= 2 {
		funcStartOffset--
	}

	return callerInfo{
		File:file,
		Line:line,
		Package:strings.Join(parts[:funcStartOffset], "."),
		Func:strings.Join(parts[funcStartOffset:], "."),
	}
}