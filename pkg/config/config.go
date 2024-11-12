package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/xmh1011/express-delivery/pkg/log"
	"github.com/xmh1011/express-delivery/pkg/storage"
	"github.com/xmh1011/express-delivery/pkg/variable"
)

// Config 定义总配置结构体
type Config struct {
	Log     log.Option     `yaml:"log"`
	Storage storage.Option `yaml:"storage"`
}

// LoadConfig 从 YAML 文件加载配置
func LoadConfig(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType(variable.DefaultConfigType)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// 将配置文件内容解析到 Config 结构体
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// AddFlags 向 FlagSet 添加命令行标志
func (c *Config) AddFlags(fs *pflag.FlagSet) {
	c.Log.AddFlags(fs, "log")
	c.Storage.AddFlags(fs, "storage")
}
