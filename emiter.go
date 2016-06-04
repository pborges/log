package log

type LogEmitter interface {
	Emit(entry *Entry)
}