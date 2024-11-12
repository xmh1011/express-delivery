package log

import (
	"github.com/spf13/pflag"

	"github.com/xmh1011/express-delivery/pkg/variable"
)

type Option struct {
	Level int `yaml:"level"`
}

func (o *Option) AddFlags(fs *pflag.FlagSet, name string) {
	flagPrefix := variable.Log
	if name != "" {
		flagPrefix += "." + name
	}

	fs.IntVar(&o.Level, flagPrefix+".level", o.Level, "log level.")
}
