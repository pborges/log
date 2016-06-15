package log

var Log *Logger

func init() {
	Log = New()
}

func NewEntry() (*Entry) {
	return Log.NewEntry()
}

func Debug(args... interface{}) {
	Log.Debug(args...)
}

func Info(args... interface{}) {
	Log.Info(args...)
}

func Warn(args... interface{}) {
	Log.Warn(args...)
}

func Error(args... interface{}) {
	Log.Error(args...)
}

func Panic(args... interface{}) {
	Log.Panic(args...)
}

func WithError(err error) (*Entry) {
	entry := Log.createEntry(2)
	return entry.WithError(err)
}

func WithField(key string, value interface{}) (*Entry) {
	entry := Log.createEntry(2)
	return entry.WithField(key, value)
}

func SetCallDepth(depth int) (*Entry) {
	entry := Log.createEntry(2 + depth)
	return entry
}