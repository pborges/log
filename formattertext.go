package log

import (
	"fmt"
	"strings"
	"strconv"
)

func NewTextFormatter() *TextFormatter {
	t := new(TextFormatter)
	t.DefaultColor = "\033[1;37m"
	t.UseColor = true
	t.ColorMap = make(map[Level]string)
	t.ColorMap[LevelDebug] = "\033[0;34m"
	t.ColorMap[LevelInfo] = "\033[0;32m"
	t.ColorMap[LevelWarn] = "\033[0;33m"
	t.ColorMap[LevelError] = "\033[1;31m"
	t.ColorMap[LevelPanic] = "\033[0;31m"
	return t
}

type TextFormatter struct {
	UseColor     bool
	DefaultColor string
	ColorMap     map[Level]string
}

func (this *TextFormatter)Format(entry Entry) (string) {
	var defaultColor, color string
	if this.UseColor {
		color = this.ColorMap[entry.Level]
		defaultColor = this.DefaultColor
	}
	entry.prependField("msg", entry.Msg)
	entry.WithField("package", entry.Package)
	entry.WithField("func", entry.Func)
	entry.WithField("file", entry.Filename + ":" + strconv.Itoa(entry.Line))

	var fields []string
	for i := 0; i < len(entry.Keys); i++ {
		fields = append(fields, fmt.Sprintf("%s%s->[ %s%+v %s]",
			defaultColor,
			entry.Keys[i],
			color,
			entry.Values[i],
			defaultColor,
		))
	}
	var fieldString = strings.Join(fields, ", ")
	return fmt.Sprintf("%s%s %s[%-5s] %-s%s\n",
		defaultColor,
		entry.TimeStamp.Format(TimeFormat),
		color,
		entry.Level.String,
		fieldString,
		defaultColor,
	)
}