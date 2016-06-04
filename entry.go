package log

import (
	"time"
	"fmt"
	"os"
)

type Entry struct {
	Msg       string
	Err       error
	log       *Logger
	CallerInfo
	Level     Level
	TimeStamp time.Time
	Keys      []string
	Values    []interface{}
}

func (this *Entry)Debug(args... interface{}) {
	this.Level = LevelDebug
	this.prepMsg(args...)
	this.log.emit(this)
}

func (this *Entry)Info(args... interface{}) {
	this.Level = LevelInfo
	this.prepMsg(args...)
	this.log.emit(this)
}

func (this *Entry)Warn(args... interface{}) {
	this.Level = LevelWarn
	this.prepMsg(args...)
	this.log.emit(this)
}

func (this *Entry)Error(args... interface{}) {
	this.Level = LevelError
	this.prepMsg(args...)
	this.log.emit(this)
}

func (this *Entry)Panic(args... interface{}) {
	this.Level = LevelPanic
	this.prepMsg(args...)
	this.log.emit(this)
	os.Exit(0)
}

func (this *Entry)WithError(err error) (*Entry) {
	this.Err = err
	return this
}

func (this *Entry)WithField(key string, value interface{}) (*Entry) {
	this.Keys = append(this.Keys, key)
	this.Values = append(this.Values, value)
	return this
}

func (this *Entry)prepMsg(args... interface{}) {
	this.Msg = fmt.Sprintln(args...)
	if len(this.Msg) >= 1 {
		this.Msg = this.Msg[:len(this.Msg) - 1]
	}
}