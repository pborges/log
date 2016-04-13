package log

import (
	"os"
	"io"
	"sync"
)

const TimeFormat string = "2006-01-02 15:04:05"

var logLevel Level = LevelDebug
var packageConfig map[string]*PackageConfig

type Logger struct {
	sync.Mutex
	Out io.Writer
}

var Formatter EntryFormatter

var logger *Logger
var entryCallDepth = 2
var hooks []HookFunc

func init() {
	hooks = make([]HookFunc, 0)
	logger = new(Logger)
	logger.Out = os.Stdout
	Formatter = NewTextFormatter()
	packageConfig = make(map[string]*PackageConfig)
}

func SetOutput(writer io.Writer) {
	logger.Out = writer
}

func SetCallDepth(depth int) (*Entry) {
	e := createEntry(logger, depth + entryCallDepth)
	return e
}

func WithField(key string, value interface{}) (*Entry) {
	e := createEntry(logger, entryCallDepth)
	e.WithField(key, value)
	return e
}

func WithError(err error) (*Entry) {
	e := createEntry(logger, entryCallDepth)
	e.WithError(err)
	return e
}

func Debug(args... interface{}) {
	e := createEntry(logger, entryCallDepth)
	e.Debug(args...)
}

func Info(args... interface{}) {
	e := createEntry(logger, entryCallDepth)
	e.Info(args...)
}

func Warn(args... interface{}) {
	e := createEntry(logger, entryCallDepth)
	e.Warn(args...)
}

func Error(args... interface{}) {
	e := createEntry(logger, entryCallDepth)
	e.Error(args...)
}

func Panic(args... interface{}) {
	e := createEntry(logger, entryCallDepth)
	e.Panic(args...)
}

func SetLogLevel(l Level) {
	logLevel = l
}

func AddHook(hook HookFunc) {
	hooks = append(hooks, hook)
}

func GetPackageConfig(pkg string) *PackageConfig {
	if conf, ok := packageConfig[pkg]; ok {
		return conf
	}
	conf := newPackageConfig()
	packageConfig[pkg] = conf
	return conf
}

func AddPermenantField(key string, value string) {
	callerInfo := getCallerInfo(1)
	GetPackageConfig(callerInfo.Package).PermanentFields[key] = value
}