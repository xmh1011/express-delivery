// cmd/insert_test.go

package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xmh1011/express-delivery/model"

	"github.com/xmh1011/express-delivery/pkg/log"
	"github.com/xmh1011/express-delivery/pkg/storage"
	"github.com/xmh1011/express-delivery/service"
)

// setupTestDatabase initializes the test database for inserting data
func setupTestDatabase() (*storage.SQLite, error) {
	// Use test config file
	configFile = setupConfigFilePath()

	// Load configuration and initialize the SQLite database
	return initConfig()
}

func TestInsertDataCommand(t *testing.T) {
	// Initialize logger
	log.InitLogger()

	// Set user and order counts
	userCount = 10
	orderCount = 50

	// Setup test database
	sqliteStorage, err := setupTestDatabase()
	assert.NoError(t, err, "Failed to initialize test database")
	assert.NotNil(t, sqliteStorage, "Expected non-nil SQLite storage")

	// Ensure cleanup of test database at the end of the test
	defer func() {
		db, _ := sqliteStorage.Client.DB()
		db.Close()
		os.Remove(filepath.Join("..", "testdata", "test.db"))
	}()

	// Migrate the Order model
	db := sqliteStorage.Client
	err = db.AutoMigrate(&model.Order{})
	assert.NoError(t, err, "Failed to migrate Order model")

	// Insert test data
	service.InsertData(db, userCount, orderCount)

	// Verify the number of inserted orders
	var orderCountInDB int64
	db.Model(&model.Order{}).Count(&orderCountInDB)
	assert.Equal(t, orderCount, orderCountInDB, "Expected number of orders in database to match orderCount")

	// Verify the number of inserted users (distinct user IDs)
	var userCountInDB int64
	db.Model(&model.Order{}).Distinct("uid").Count(&userCountInDB)
	assert.Equal(t, userCount, userCountInDB, "Expected number of distinct users in database to match userCount")
}
