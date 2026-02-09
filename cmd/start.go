package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start subject of subcommand",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please specify the subject to be started")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
