package log

import (
	"time"
	"sync"
)

func New() (l *Logger) {
	l = new(Logger)
	l.Emitter = NewEmitterText()
	l.Mutex = new(sync.Mutex)
	return
}

type Logger struct {
	*sync.Mutex
	Emitter LogEmitter
	Level   Level
}

func (this *Logger)NewEntry() (*Entry) {
	return Log.createEntry(3)
}

func (this *Logger)Debug(args... interface{}) {
	entry := this.createEntry(3)
	entry.Debug(args...)
}

func (this *Logger)Info(args... interface{}) {
	entry := this.createEntry(3)
	entry.Info(args...)
}

func (this *Logger)Warn(args... interface{}) {
	entry := this.createEntry(3)
	entry.Warn(args...)
}

func (this *Logger)Error(args... interface{}) {
	entry := this.createEntry(3)
	entry.Error(args...)
}

func (this *Logger)Panic(args... interface{}) {
	entry := this.createEntry(3)
	entry.Panic(args...)
}

func (this *Logger)WithError(err error) *Entry {
	entry := this.createEntry(3)
	entry.WithError(err)
	return entry
}

func (this *Logger)WithField(key string, value interface{}) *Entry {
	entry := this.createEntry(3)
	entry.WithField(key, value)
	return entry
}

func (this *Logger)createEntry(callSkip int) (*Entry) {
	e := new(Entry)
	e.log = this
	e.CallerInfo = getCallerInfo(callSkip)
	e.TimeStamp = time.Now()
	return e
}

func (this *Logger)emit(entry *Entry) {
	if entry.Level.Priority >= this.Level.Priority {
		this.Lock()
		defer this.Unlock()
		this.Emitter.Emit(entry)
	}
}