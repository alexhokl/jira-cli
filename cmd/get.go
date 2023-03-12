package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get subject of subcommand",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please specify the subject to be retrieved.")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
