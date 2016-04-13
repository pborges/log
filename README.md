#Because what the world needs is another logger.
I love logrus, but I want line and file numbers, and for whatever reason logrus could not supply that.

Here is a potentially very bad logger, somewhat interoperable with logrus

## Features

*   Color Output default
*   EntryFormatter interface for customizing log output
*   Included JSON formatter
*   Include fields in logging
*   Permanent Fields by package name
*   Set LogLevel by package name
*   Simple hooking system
*   Adds Caller info (package, function name, line number, etc)
*   Dynamically set call depth for use in generic logging functions
*   Probably more

I am not married to what the output looks like, but I needed to pad (atleast the filename/linenumber) with spaces so intellij would recognize it and make it a hyperlink

```
package main

import (
	"github.com/pborges/log"
	"fmt"
	"time"
)

func main() {
	log.AddHook(func(entry log.Entry) {
		if entry.Level.Priority > log.LevelWarn.Priority {
			fmt.Println("OMG WARNING")
		}
	})

	log.Debug("hello", "world")

	data := struct {
		String string
		Int    int
	}{
		"string",
		1,
	}

	// log a struct
	log.WithField("d", data).Debug("some data")

	// notice the line number in this log output
	unifiedLoggingFunc("blargh")

	log.Debug("before debug")
	log.SetLogLevel(log.LevelInfo)
	log.Debug("after debug")
	log.Info("after info")

	log.WithField("animal", "not a walrus").Info("some message")
	log.WithField("animal", "not a walrus").Warn("some message")
	log.WithField("animal", "not a walrus").Error("some message")

	time.Sleep(1 * time.Second)
	log.Panic("ZOMG")
}

func unifiedLoggingFunc(message string) {
	log.SetCallDepth(1).Debug(message)
}
```

![Alt text](screenshot/output.png?raw=true)
