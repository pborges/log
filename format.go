package log

type EntryFormatter interface {
	Format(entry Entry) (string)
}
