#Because what the world needs is another logger.

##Disclaimer: this is my personal logger, your welcome to fork but dont rely on this repo as it as there is no api guarantee

I love logrus, but I want line and file numbers, and for whatever reason logrus could not supply that.

Here is a potentially very bad logger, somewhat interoperable with logrus

## Features

*   Color Output default
*   LogEmitter interface for customizing log output
*   Include fields in logging
*   Adds Caller info (package, function name, line number, etc)
*   Dynamically set call depth for use in generic logging functions
*   Probably more

```
package main

import (
	"github.com/pborges/log"
	"fmt"
)

func main() {
	log.Debug("hello", "world")
	log.WithField("test", 3).WithField("fsdaf", 5).Info("hello2", "world2")

	data := struct {
		NamedString string
		NamedInt    int
	}{
		"string",
		1,
	}

	e := StructuredError{"blah", 65}
	log.WithError(e).Warn("ERROR TEST")

	//
	// log a struct
	log.WithField("d", data).Debug("some data")

	// notice the line number in this log output
	unifiedLoggingFunc("blargh")
	log.Debug("before debug")
	log.Log.Level = log.LevelInfo
	log.Debug("after debug")
	log.Info("after info")

	log.WithField("animal", "not a walrus").Info("some message")
	log.WithField("animal", "not a walrus").WithField("number_of_legs", 4).Warn("some message")
	log.WithField("animal", "not a walrus").Error("some message")
	log.Panic("ZOMG")
}

func unifiedLoggingFunc(message string) {
	log.SetCallDepth(1).Debug(message)
}

type StructuredError struct {
	Data1 string
	Data2 int
}

func (this StructuredError)Error() string {
	return fmt.Sprintf("error : %s %d", this.Data1, this.Data2)
}
```

![Alt text](screenshot/output.png?raw=true)
