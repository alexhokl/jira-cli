package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alexhokl/jira-cli/swagger_software"
)

type startSprintOptions struct {
	id int64
}

var startSprintOpts = startSprintOptions{}

var startSprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Start a sprint",
	RunE:  runStartSprint,
}

func init() {
	startCmd.AddCommand(startSprintCmd)

	flags := startSprintCmd.Flags()
	flags.Int64VarP(&startSprintOpts.id, "id", "i", 0, "Sprint ID (required)")

	startSprintCmd.MarkFlagRequired("id")
}

func runStartSprint(_ *cobra.Command, _ []string) error {
	request := swagger_software.NewUpdateSprintRequest()
	request.SetState("active")

	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	_, err := client.SprintAPI.PartiallyUpdateSprint(ctx, startSprintOpts.id).UpdateSprintRequest(*request).Execute()
	if err != nil {
		return wrapAPIError(err)
	}

	fmt.Println("Sprint started")
	return nil
}
