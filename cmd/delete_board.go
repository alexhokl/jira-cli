package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type deleteBoardOptions struct {
	boardId int64
}

var deleteBoardOpts = deleteBoardOptions{}

var deleteBoardCmd = &cobra.Command{
	Use:   "board",
	Short: "Delete a board",
	Long: `Delete a board by its ID.

Note: Deleting a board does not delete the issues on the board.
The board ID can be found using 'jira-cli list boards'.

Examples:
  # Delete a board
  jira-cli delete board --id 123`,
	RunE: runDeleteBoard,
}

func init() {
	deleteCmd.AddCommand(deleteBoardCmd)

	flags := deleteBoardCmd.Flags()
	flags.Int64VarP(&deleteBoardOpts.boardId, "id", "i", 0, "Board ID (required)")
	deleteBoardCmd.MarkFlagRequired("id")
}

func runDeleteBoard(_ *cobra.Command, _ []string) error {
	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	_, err := client.BoardAPI.DeleteBoard(ctx, deleteBoardOpts.boardId).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	fmt.Printf("Board %d has been deleted\n", deleteBoardOpts.boardId)

	return nil
}
