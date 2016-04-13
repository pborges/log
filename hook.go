package log

type HookFunc func(entry Entry)

func processHooks(entry Entry) {
	for _, hook := range hooks {
		hook(entry)
	}
}