package log

import (
	"fmt"
	"html/template"
	"io"
	"os"
)

func NewEmitterText() (e *EmitterText) {
	e = new(EmitterText)
	e.Output = os.Stdout
	fm := template.FuncMap{}
	fm["fmt_pad"] = func(str string, length int) string {
		for i := len(str); i < length; i++ {
			str += " "
		}
		return str
	}

	fm["fmt_idx"] = func(arr []interface{}, idx int) string {
		return fmt.Sprintf("%+v", arr[idx])
	}

	var err error
	e.Template, err = template.New("").Funcs(fm).ParseFiles("log.tmpl")
	if err != nil {
		panic(err)
	}
	return
}

type EmitterText struct {
	Template *template.Template
	Output   io.Writer
}

func (this *EmitterText)Emit(entry *Entry) {
	model := struct {
		*Entry
		ColorDefault   string
		ColorSubtle    string
		ColorHighlight string
		ColorError     string
	}{
		Entry:entry,
		ColorDefault:"\033[1;37m",
		ColorSubtle:"\033[1;30m",
		ColorError:"\033[1;31m",
	}

	switch entry.Level{
	case LevelDebug:
		model.ColorHighlight = "\033[0;34m"
	case LevelInfo:
		model.ColorHighlight = "\033[0;32m"
	case LevelWarn:
		model.ColorHighlight = "\033[0;33m"
	case LevelError:
		model.ColorHighlight = "\033[1;31m"
	case LevelPanic:
		model.ColorHighlight = "\033[0;31m"
	}

	err := this.Template.ExecuteTemplate(this.Output, "log.tmpl", model)
	if err != nil {
		panic(err)
	}
}