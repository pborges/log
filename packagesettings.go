package log

func newPackageConfig() *PackageConfig {
	p := PackageConfig{}
	p.PermanentFields = make(map[string]interface{})
	p.Level = LevelInfo
	return &p
}

type PackageConfig struct {
	Level           Level
	PermanentFields map[string]interface{}
}