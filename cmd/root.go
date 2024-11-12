// cmd/root.go

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "express-delivery",
	Short: "Express delivery application to manage orders and calculate costs",
	Long: `An application to manage express delivery orders, calculate shipping costs,
insert sample data, query user order details, and display application version.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(insertCmd)
	rootCmd.AddCommand(queryCmd)
}
