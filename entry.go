package log

import (
	"fmt"
	"time"
	"encoding/json"
)

func createEntry(logger *Logger, callSkip int) (*Entry) {
	e := new(Entry)
	e.Keys = make([]string, 0)
	e.Values = make([]interface{}, 0)

	callInfo := getCallerInfo(callSkip)

	e.Package = callInfo.Package
	e.Func = callInfo.Func
	e.Filename = callInfo.File
	e.Line = callInfo.Line

	e.TimeStamp = time.Now()

	if conf, ok := packageConfig[e.Package]; ok {
		for k, v := range conf.PermanentFields {
			e.WithField(k, v)
		}
	}
	return e
}

type Entry struct {
	Level     Level
	Msg       string
	Filename  string
	Package   string
	Func      string
	Line      int
	Keys      []string      // don't use a hashmap to preserve order
	Values    []interface{} // don't use a hashmap to preserve order
	TimeStamp time.Time
}

func (entry *Entry) MarshalJSON() ([]byte, error) {
	jsonEntry := struct {
		Timestamp time.Time `json:"timestamp"`
		Msg       string `json:"msg"`
		FileName  string `json:"filename"`
		Package   string `json:"package"`
		Func      string `json:"func"`
		Line      int `json:"line"`
		Level     string `json:"level"`
		Fields    map[string]interface{} `json:"fields"`
	}{
		Msg:entry.Msg,
		FileName:entry.Filename,
		Package:entry.Package,
		Line:entry.Line,
		Func:entry.Func,
		Level:entry.Level.String,
		Timestamp:entry.TimeStamp,
		Fields:make(map[string]interface{}),
	}

	for i, key := range entry.Keys {
		jsonEntry.Fields[key] = entry.Values[i]
	}

	return json.Marshal(jsonEntry) // how do you handle errors in a log package? ha
}

func (this *Entry)prepMsg(args... interface{}) {
	this.Msg = fmt.Sprintln(args...)
	if len(this.Msg) >= 1 {
		this.Msg = this.Msg[:len(this.Msg) - 1]
	}
}

func (this *Entry)emit(level Level, args... interface{}) {
	this.Level = level
	this.prepMsg(args...)
	processHooks(*this)
	logger.Lock()
	defer logger.Unlock()
	logger.Out.Write([]byte(Formatter.Format(*this)))
}

func (this *Entry)prependField(key string, value interface{}) (*Entry) {
	this.Keys = append([]string{key}, this.Keys...)
	this.Values = append([]interface{}{value}, this.Values...)
	return this
}

func (this *Entry)WithField(key string, value interface{}) (*Entry) {
	this.Keys = append(this.Keys, key)
	this.Values = append(this.Values, value)
	return this
}

func (this *Entry)WithError(err error) (*Entry) {
	this.WithField("err", err)
	return this
}