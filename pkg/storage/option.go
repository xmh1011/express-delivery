package storage

import (
	"github.com/spf13/pflag"

	"github.com/xmh1011/express-delivery/pkg/variable"
)

// Option defines options for sqlite database.
type Option struct {
	SourceName string `yaml:"sourceName,omitempty"` // 数据库文件名
	LogLevel   int    `yaml:"log-level,omitempty"`
}

func (o *Option) AddFlags(fs *pflag.FlagSet, name string) {
	flagPrefix := variable.DBEngineSQLite
	if name != "" {
		flagPrefix += "." + name
	}

	fs.StringVar(&o.SourceName, flagPrefix+".source-name", o.SourceName, "SQLite source name.")
}
