package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/xmh1011/express-delivery/pkg/config"
	"github.com/xmh1011/express-delivery/pkg/log"
	"github.com/xmh1011/express-delivery/pkg/storage"
)

var configFile string
var cfg *config.Config

const (
	configFlagShort = "c"
	configFlag      = "config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Load and display configuration",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, err = config.LoadConfig(configFile)
		if err != nil {
			log.InitLogger().Fatalf("Error loading config file: %v", err)
		}

		// 根据配置文件中的日志级别重新初始化 logger
		logger := log.InitLogger()
		logger.SetLevel(logrus.Level(cfg.Log.Level)) // 使用配置中的日志级别
		logger.Debugf("Loaded Configuration:\n")
		logger.Debugf("Log Level: %d\n", cfg.Log.Level)
		logger.Debugf("SQLite Source Name: %s\n", cfg.Storage.SourceName)
		logger.Debugf("SQLite Log Level: %d\n", cfg.Storage.LogLevel)
	},
}

// initConfig 加载配置文件并初始化 logger 和数据库
func initConfig() (*storage.SQLite, error) {
	var err error
	cfg, err = config.LoadConfig(configFile)
	if err != nil {
		return nil, err
	}

	// 根据配置文件中的日志级别初始化 logger
	logger := log.InitLogger()
	logger.SetLevel(logrus.Level(cfg.Log.Level)) // 使用配置中的日志级别

	// 初始化 SQLite 数据库
	sqliteStorage := &storage.SQLite{}
	err = sqliteStorage.Init(&storage.Option{
		SourceName: cfg.Storage.SourceName,
		LogLevel:   cfg.Storage.LogLevel,
	})
	if err != nil {
		return nil, err
	}

	return sqliteStorage, nil
}

func init() {
	// 定义 config 命令的标志
	configCmd.PersistentFlags().StringVarP(&configFile, configFlag, configFlagShort, "", "Path to the config file")
}
