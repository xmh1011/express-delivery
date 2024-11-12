package storage

// Option defines options for sqlite database.
type Option struct {
	SourceName string `yaml:"sourceName,omitempty"` // 数据库文件名
	LogLevel   int    `yaml:"log-level,omitempty"`
}
