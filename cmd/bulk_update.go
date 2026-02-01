package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bulkUpdateCmd = &cobra.Command{
	Use:   "bulk-update",
	Short: "Bulk update operations",
	Long:  `Perform bulk update operations on multiple issues.`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please specify a bulk-update subcommand (e.g., rename-label)")
	},
}

func init() {
	rootCmd.AddCommand(bulkUpdateCmd)
}
