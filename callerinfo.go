package log

import (
	"runtime"
	"strings"
)

type callerInfo struct {
	File    string
	Line    int
	Package string
	Func    string
}

func getCallerInfo(callSkip int) callerInfo {
	pc, file, line, _ := runtime.Caller(callSkip + 1) // plus one is to account for itself
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl - 1]

	if parts[pl - 2][0] == '(' {
		funcName = parts[pl - 2] + "." + funcName
		packageName = strings.Join(parts[0:pl - 2], ".")
	} else {
		packageName = strings.Join(parts[0:pl - 1], ".")
	}
	return callerInfo{
		File:file,
		Line:line,
		Package:packageName,
		Func:funcName,
	}
}