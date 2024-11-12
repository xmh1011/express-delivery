// cmd/query.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xmh1011/express-delivery/pkg/log"
	"github.com/xmh1011/express-delivery/service"
)

const (
	uidFlagShort = "u"
	uidFlag      = "uid"
)

var uid int64

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the express delivery cost for a given user",
	Run: func(cmd *cobra.Command, args []string) {
		// 加载配置并初始化 logger 和数据库
		sqliteStorage, err := initConfig()
		if err != nil {
			log.InitLogger().Fatalf("Error loading config file: %v", err)
		}
		db := sqliteStorage.Client

		// create logger
		logger := log.InitLogger()
		logger.Infof("Querying total delivery cost for user %d", uid)

		// 查询用户订单并计算总费用
		totalCost, err := service.QueryUserOrders(db, uid)
		if err != nil {
			logger.Fatalf("Error querying orders: %v", err)
		}

		fmt.Printf("Total delivery cost for user %d: %.2f\n", uid, totalCost)
	},
}

func init() {
	queryCmd.Flags().Int64VarP(&uid, uidFlag, uidFlagShort, 1, "The uid of the user (required)")
	queryCmd.PersistentFlags().StringVarP(&configFile, configFlag, configFlagShort, "", "Path to the config file")
}
