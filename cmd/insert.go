// cmd/insert.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xmh1011/express-delivery/model"
	"github.com/xmh1011/express-delivery/pkg/log"
	"github.com/xmh1011/express-delivery/service"
)

var userCount int64
var orderCount int64

const (
	userCountShort  = "u"
	userCountFlag   = "user"
	orderCountShort = "o"
	orderCountFlag  = "order"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert generated user and order data into the database",
	Run: func(cmd *cobra.Command, args []string) {
		// 加载配置并初始化 logger 和数据库
		sqliteStorage, err := initConfig()
		if err != nil {
			log.InitLogger().Fatalf("Error loading config file: %v", err)
		}
		db := sqliteStorage.Client

		// create logger
		logger := log.InitLogger()
		logger.Infof("Inserting data: %d users and %d orders", userCount, orderCount)

		// 迁移数据库模型并插入数据
		err = db.AutoMigrate(&model.Order{})
		if err != nil {
			logger.Fatalf("Failed to migrate Order model: %v", err)
		}
		service.InsertData(db, userCount, orderCount)

		fmt.Printf("Successfully inserted %d users and %d orders into the database.\n", userCount, orderCount)
	},
}

func init() {
	insertCmd.Flags().Int64VarP(&userCount, userCountFlag, userCountShort, 1000, "Number of users to generate")
	insertCmd.Flags().Int64VarP(&orderCount, orderCountFlag, orderCountShort, 100000, "Number of orders to generate")
	insertCmd.PersistentFlags().StringVarP(&configFile, configFlag, configFlagShort, "", "Path to the config file")
}
