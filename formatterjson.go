package log

type JsonFormatter struct{}

func (this *JsonFormatter)Format(entry Entry) (string) {
	j, err := entry.MarshalJSON() // how do you handle errors in a log package? ha
	if err != nil {
		panic(err)
	}
	return string(j) + "\n"
}