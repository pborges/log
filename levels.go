package log

type Level struct {
	String   string
	Priority int
}

var LevelDebug Level = Level{"DEBUG", 0}
var LevelInfo Level = Level{"INFO", 1}
var LevelWarn Level = Level{"WARN", 2}
var LevelError Level = Level{"ERROR", 3}
var LevelPanic Level = Level{"PANIC", 4}