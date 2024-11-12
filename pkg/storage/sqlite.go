package storage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/xmh1011/express-delivery/pkg/log"
)

type SQLite struct {
	Config *Option
	Client *gorm.DB
}

func (m *SQLite) Init(opt *Option) error {
	m.Config = opt

	// SQLite 数据源名称通常是数据库文件路径，例如 "test.db" 或 ":memory:" 用于内存数据库
	dsn := opt.SourceName
	log.L.Debugf("SQLite connection dsn: %s", dsn)

	// 使用 SQLite 驱动和提供的日志级别打开数据库
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(opt.LogLevel)),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to SQLite database: %w", err)
	}

	m.Client = db
	return nil
}
