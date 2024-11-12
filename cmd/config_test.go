// cmd/config_test.go

package cmd

import (
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/xmh1011/express-delivery/pkg/config"
	"github.com/xmh1011/express-delivery/pkg/log"
)

// 设置测试配置文件路径
func setupConfigFilePath() string {
	return filepath.Join(".", "testdata", "config.yaml")
}

func TestConfigCmdLoadConfig(t *testing.T) {
	// 指定配置文件路径
	configFile = setupConfigFilePath()

	// 加载配置
	var err error
	cfg, err = config.LoadConfig(configFile)
	assert.NoError(t, err, "Error loading config file")
	assert.NotNil(t, cfg, "Expected non-nil configuration")

	// 检查配置的正确性
	assert.Equal(t, 5, cfg.Log.Level, "Expected log level to be 5 from config file")
	assert.Equal(t, "test.db", cfg.Storage.SourceName, "Expected SQLite source name to be 'test.db'")
	assert.Equal(t, 1, cfg.Storage.LogLevel, "Expected SQLite log level to be 1 from config file")
}

func TestInitConfig(t *testing.T) {
	// 指定配置文件路径
	configFile = setupConfigFilePath()

	// 初始化配置、日志和数据库
	sqliteStorage, err := initConfig()
	assert.NoError(t, err, "Error initializing config, logger, and SQLite")
	assert.NotNil(t, sqliteStorage, "Expected non-nil SQLite storage")

	// 验证日志级别是否正确设置
	logger := log.InitLogger()
	assert.Equal(t, logrus.Level(cfg.Log.Level), logger.GetLevel(), "Expected logger level to match config level")

	// 检查数据库连接是否正常
	db, err := sqliteStorage.Client.DB()
	assert.NoError(t, err, "Failed to get generic database object")
	assert.NoError(t, db.Ping(), "Expected successful ping to the SQLite database")

	// 测试结束后关闭数据库连接
	db.Close()
}
