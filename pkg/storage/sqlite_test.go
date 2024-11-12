// storage_test.go
package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/xmh1011/express-delivery/model"
	"github.com/xmh1011/express-delivery/pkg/log"
)

// 初始化表结构的辅助函数
func initializeTestDatabase(db *gorm.DB) error {
	return db.AutoMigrate(&model.Order{})
}

func TestSQLiteInitWithFileDatabase(t *testing.T) {
	// 初始化日志
	log.InitLogger()

	// 指定数据库文件路径
	dbPath := filepath.Join("testdata", "test.db")
	opt := &Option{
		SourceName: dbPath, // 使用文件数据库
		LogLevel:   1,      // 设置日志级别
	}

	// 检查文件是否存在，若不存在则创建表结构
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		// 初始化 SQLite 实例并创建表结构
		sqlite := &SQLite{}
		err = sqlite.Init(opt)
		if err != nil {
			t.Fatalf("Failed to initialize SQLite: %v", err)
		}
		err = initializeTestDatabase(sqlite.Client)
		if err != nil {
			t.Fatalf("Failed to initialize test database structure: %v", err)
		}
		// 关闭数据库
		db, _ := sqlite.Client.DB()
		db.Close()
	}

	// 测试数据库连接
	sqlite := &SQLite{}
	err := sqlite.Init(opt)

	// 确认没有错误发生
	assert.NoError(t, err, "Failed to initialize SQLite with file database")

	// 确认数据库连接成功
	assert.NotNil(t, sqlite.Client, "Expected a non-nil *gorm.DB client")

	// 确认数据库连接可用
	db, err := sqlite.Client.DB()
	assert.NoError(t, err, "Failed to get generic database object")
	assert.NoError(t, db.Ping(), "Expected successful ping to the SQLite file database")

	// 测试结束后关闭数据库连接
	db.Close()
}
