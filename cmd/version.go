package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xmh1011/express-delivery/pkg/variable"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(formatVersionInfo())
	},
}

// formatVersionInfo returns the application version and Git commit info as a formatted string
func formatVersionInfo() string {
	return fmt.Sprintf("Application Version: %s\nGit Commit Message: %s", variable.Version, variable.GitCommit)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
