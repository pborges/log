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

func (this *Entry)Debug(args... interface{}) {
	if logPriorityCheck(LevelDebug, *this) {
		this.emit(LevelDebug, args...)
	}
}

func (this *Entry)Info(args... interface{}) {
	if logPriorityCheck(LevelInfo, *this) {
		this.emit(LevelInfo, args...)
	}
}

func (this *Entry)Warn(args... interface{}) {
	if logPriorityCheck(LevelWarn, *this) {
		this.emit(LevelWarn, args...)
	}
}

func (this *Entry)Error(args... interface{}) {
	if logPriorityCheck(LevelError, *this) {
		this.emit(LevelError, args...)
	}
}

func (this *Entry)Panic(args... interface{}) {
	this.prepMsg(args...)
	this.Level = LevelPanic
	panic(Formatter.Format(*this))
}

// logPriorityCheck will return true if this log entry should be emitted
func logPriorityCheck(level Level, e Entry) bool {
	if level.Priority < logLevel.Priority {
		return false
	}
	if p, ok := packageConfig[e.Package]; ok {
		return level.Priority >= p.Level.Priority
	}
	return true
}