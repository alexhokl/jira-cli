package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a subject of subcommand",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please specify the subject to be added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
