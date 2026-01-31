package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger_software"
)

type closeSprintOptions struct {
	id int64
}

var closeSprintOpts = closeSprintOptions{}

var closeSprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Close a sprint",
	RunE:  runCloseSprint,
}

func init() {
	closeCmd.AddCommand(closeSprintCmd)

	flags := closeSprintCmd.Flags()
	flags.Int64VarP(&closeSprintOpts.id, "id", "i", 0, "Sprint ID (required)")

	closeSprintCmd.MarkFlagRequired("id")
}

func runCloseSprint(_ *cobra.Command, _ []string) error {
	request := swagger_software.NewUpdateSprintRequest()
	request.SetState("closed")

	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	_, err := client.SprintAPI.PartiallyUpdateSprint(ctx, closeSprintOpts.id).UpdateSprintRequest(*request).Execute()
	if err != nil {
		return err
	}

	fmt.Println("Sprint closed")
	return nil
}
